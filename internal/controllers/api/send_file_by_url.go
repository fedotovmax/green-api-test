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

type sendFileByURLSignals struct {
	Token      string `json:"apiToken"`
	InstanceID string `json:"instanceId"`
	ChatID     string `json:"sendFileChatId"`
	FileURL    string `json:"fileUrl"`
	FileName   string `json:"fileName"`
}

func (c *controller) sendFileByUrl(w http.ResponseWriter, r *http.Request) {
	var signals sendFileByURLSignals

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

	input := &inputs.SendFile{ChatID: signals.ChatID, FileURL: signals.FileURL, FileName: signals.FileName}

	err = input.Validate()

	if err != nil {
		sse.ExecuteScript(fmt.Sprintf("alert(`%s`)", err.Error()))
		return
	}

	apiCtx, cancelApiCtx := context.WithTimeout(r.Context(), time.Second*5)
	defer cancelApiCtx()

	newMessage, err := c.greenApi.SendFileByUrl(
		apiCtx,
		creds.InstanceID,
		creds.APIToken,
		input,
	)

	if err != nil {
		sse.ExecuteScript(fmt.Sprintf("alert(`%s`)", err.Error()))
		return
	}

	err = sse.PatchElementTempl(
		components.NewMessage(newMessage),
		datastar.WithSelectorID(dom.GreenAPIResponseSelectorID),
		datastar.WithModeAppend(),
	)

	if err != nil {
		sse.ExecuteScript(fmt.Sprintf("alert(`%s`)", err.Error()))
		return
	}
}
