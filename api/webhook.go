//
// Copyright (C) 2020 assay.it
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/assay.it/assay
//

package api

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/assay-it/sdk-go/assay"
	ç "github.com/assay-it/sdk-go/cats"
	"github.com/assay-it/sdk-go/http"
	ƒ "github.com/assay-it/sdk-go/http/recv"
	ø "github.com/assay-it/sdk-go/http/send"
)

func stdout(hook *[]byte) assay.Arrow {
	return ç.FMap(func() error {
		var pretty bytes.Buffer
		if err := json.Indent(&pretty, *hook, "", "  "); err != nil {
			return err
		}
		fmt.Println(pretty.String())
		return nil
	})
}

//
func (c *Client) WebHookSource(req SourceCodeID) assay.Arrow {
	return c.webhook("https://%s/webhook/sourcecode", req)
}

//
func (c *Client) WebHookCommit(req SourceCodeID) assay.Arrow {
	return c.webhook("https://%s/webhook/commit", req)
}

//
func (c *Client) WebHookRelease(req SourceCodeID) assay.Arrow {
	return c.webhook("https://%s/webhook/release", req)
}

//
func (c *Client) webhook(uri string, req SourceCodeID) assay.Arrow {
	var hook []byte

	return http.Join(
		ø.POST(uri, c.api),
		ø.Authorization().Val(&c.token),
		ø.ContentJSON(),
		ø.Send(req),
		ƒ.Code(http.StatusOK),
		ƒ.Bytes(&hook),
	).Then(
		stdout(&hook),
	)
}

//
func (c *Client) WebHook(req Hook) assay.Arrow {
	var hook []byte

	return http.Join(
		ø.POST("https://%s/webhook", c.api),
		ø.Authorization().Val(&c.token),
		ø.ContentJSON(),
		ø.Send(req),
		ƒ.Code(http.StatusOK),
		ƒ.Bytes(&hook),
	).Then(
		stdout(&hook),
	)
}
