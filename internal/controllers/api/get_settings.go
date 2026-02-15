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

type getSettingsSignals struct {
	Token      string `json:"apiToken"`
	InstanceID string `json:"instanceId"`
}

func (c *controller) getSettings(w http.ResponseWriter, r *http.Request) {

	var signals getSettingsSignals

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

	settings, err := c.greenApi.GetSettings(apiCtx, signals.InstanceID, signals.Token)

	if err != nil {
		sse.ExecuteScript(fmt.Sprintf("alert(`%s`)", err.Error()))
		return
	}

	err = sse.PatchElementTempl(
		components.Settings(settings),
		datastar.WithSelectorID(dom.GreenAPIResponseSelectorID),
		datastar.WithModeAppend(),
	)

	if err != nil {
		sse.ExecuteScript(fmt.Sprintf("alert(`%s`)", err.Error()))
		return
	}

}
