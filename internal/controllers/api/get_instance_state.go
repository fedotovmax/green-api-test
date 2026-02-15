package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fedotovmax/green-api-test/internal/dom"
	"github.com/fedotovmax/green-api-test/internal/domain/inputs"
	"github.com/fedotovmax/green-api-test/internal/templates/components"
	"github.com/starfederation/datastar-go/datastar"
)

type getStateSignals struct {
	Token      string `json:"apiToken"`
	InstanceID string `json:"instanceId"`
}

func (c *controller) getStateInstance(w http.ResponseWriter, r *http.Request) {
	var signals getStateSignals

	err := datastar.ReadSignals(r, &signals)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)

	creds := inputs.Credentials{APIToken: signals.Token, InstanceID: signals.InstanceID}

	err = creds.Validate()

	if err != nil {
		sse.ExecuteScript(fmt.Sprintf("alert(`%s`)", err.Error()))
		return
	}

	apiCtx, cancelApiCtx := context.WithTimeout(r.Context(), time.Second*5)
	defer cancelApiCtx()

	state, err := c.greenApi.GetStateInstance(apiCtx, signals.InstanceID, signals.Token)

	if err != nil {
		sse.ExecuteScript(fmt.Sprintf("alert(`%s`)", err.Error()))
		return
	}

	err = sse.PatchElementTempl(
		components.State(state),
		datastar.WithSelectorID(dom.GreenAPIResponseSelectorID),
		datastar.WithModeAppend(),
	)

	if err != nil {
		sse.ExecuteScript(fmt.Sprintf("alert(`%s`)", err.Error()))
		return
	}
}
