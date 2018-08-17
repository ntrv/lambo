package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ntrv/lambo/lambo"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

func (h Hook) runProcessContext(
	ctx context.Context,
	fn lambo.HandleProcessFunc,
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	payload := []byte(req.Body)
	switch h.eventName {
	case gh.CommitCommentEvent:
		var cc gh.CommitCommentPayload
		json.Unmarshal(payload, &cc)
		return fn(ctx, cc, req)
	case gh.CreateEvent:
		var cr gh.CreatePayload
		json.Unmarshal(payload, &cr)
		return fn(ctx, cr, req)
	case gh.DeleteEvent:
		var de gh.DeletePayload
		json.Unmarshal(payload, &de)
		return fn(ctx, de, req)
	case gh.DeploymentEvent:
		var dp gh.DeploymentPayload
		json.Unmarshal(payload, &dp)
		return fn(ctx, dp, req)
	case gh.DeploymentStatusEvent:
		var ds gh.DeploymentStatusPayload
		json.Unmarshal(payload, &ds)
		return fn(ctx, ds, req)
	case gh.ForkEvent:
		var fk gh.ForkPayload
		json.Unmarshal(payload, &fk)
		return fn(ctx, fk, req)
	case gh.GollumEvent:
		var gl gh.GollumPayload
		json.Unmarshal(payload, &gl)
		return fn(ctx, gl, req)
	case gh.InstallationEvent, gh.IntegrationInstallationEvent:
		var in gh.InstallationPayload
		json.Unmarshal(payload, &in)
		return fn(ctx, in, req)
	case gh.IssueCommentEvent:
		var ic gh.IssueCommentPayload
		json.Unmarshal(payload, &ic)
		return fn(ctx, ic, req)
	case gh.IssuesEvent:
		var is gh.IssuesPayload
		json.Unmarshal(payload, &is)
		return fn(ctx, is, req)
	case gh.LabelEvent:
		var lb gh.LabelPayload
		json.Unmarshal(payload, &lb)
		return fn(ctx, lb, req)
	case gh.MemberEvent:
		var me gh.MemberPayload
		json.Unmarshal(payload, &me)
		return fn(ctx, me, req)
	case gh.MembershipEvent:
		var ms gh.MembershipPayload
		json.Unmarshal(payload, &ms)
		return fn(ctx, ms, req)
	case gh.MilestoneEvent:
		var mi gh.MilestonePayload
		json.Unmarshal(payload, &mi)
		return fn(ctx, mi, req)
	case gh.OrganizationEvent:
		var or gh.OrganizationPayload
		json.Unmarshal(payload, &or)
		return fn(ctx, or, req)
	case gh.OrgBlockEvent:
		var ob gh.OrgBlockPayload
		json.Unmarshal(payload, &ob)
		return fn(ctx, ob, req)
	case gh.PageBuildEvent:
		var pa gh.PageBuildPayload
		json.Unmarshal(payload, &pa)
		return fn(ctx, pa, req)
	case gh.PingEvent:
		var pi gh.PingPayload
		json.Unmarshal(payload, &pi)
		return fn(ctx, pi, req)
	case gh.ProjectCardEvent:
		var pc gh.ProjectCardPayload
		json.Unmarshal(payload, &pc)
		return fn(ctx, pc, req)
	case gh.ProjectColumnEvent:
		var po gh.ProjectColumnPayload
		json.Unmarshal(payload, &po)
		return fn(ctx, po, req)
	case gh.ProjectEvent:
		var pe gh.ProjectPayload
		json.Unmarshal(payload, &pe)
		return fn(ctx, pe, req)
	case gh.PublicEvent:
		var pu gh.PublicPayload
		json.Unmarshal(payload, &pu)
		return fn(ctx, pu, req)
	case gh.PullRequestEvent:
		var pr gh.PullRequestPayload
		json.Unmarshal(payload, &pr)
		return fn(ctx, pr, req)
	case gh.PullRequestReviewEvent:
		var prr gh.PullRequestReviewPayload
		json.Unmarshal(payload, &prr)
		return fn(ctx, prr, req)
	case gh.PullRequestReviewCommentEvent:
		var prc gh.PullRequestReviewCommentPayload
		json.Unmarshal(payload, &prc)
		return fn(ctx, prc, req)
	case gh.PushEvent:
		var pu gh.PushPayload
		json.Unmarshal(payload, &pu)
		return fn(ctx, pu, req)
	case gh.ReleaseEvent:
		var re gh.ReleasePayload
		json.Unmarshal(payload, &re)
		return fn(ctx, re, req)
	case gh.RepositoryEvent:
		var rp gh.RepositoryPayload
		json.Unmarshal(payload, &rp)
		return fn(ctx, rp, req)
	case gh.StatusEvent:
		var st gh.StatusPayload
		json.Unmarshal(payload, &st)
		return fn(ctx, st, req)
	case gh.TeamEvent:
		var te gh.TeamPayload
		json.Unmarshal(payload, &te)
		return fn(ctx, te, req)
	case gh.TeamAddEvent:
		var ta gh.TeamAddPayload
		json.Unmarshal(payload, &ta)
		return fn(ctx, ta, req)
	case gh.WatchEvent:
		var wa gh.WatchPayload
		json.Unmarshal(payload, &wa)
		return fn(ctx, wa, req)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
	}, fmt.Errorf("")
}
