package localize

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

var (
	pageMsgs = map[string][]string{
		"notfound.go.html": {
			"Page404Title",
			"Page404Description",
			"Page404Header",
			"Page404Message",
			"Page404Button",
		},
		"view_note.go.html": {
			"PageViewNoteKeywords",
			"PageViewNoteCopyLink",
			"PageViewNoteShare",
			"PageViewNoteEdit",
			"PageViewNoteReport",
			"PageViewNoteEnterPassword",
			"PageViewNotePassword",
			"PageViewNoteEditNote",
			"PageViewNoteReason",
			"PageViewNoteSelectAReason",
			"PageViewNotePornAdultContent",
			"PageViewNoteCopyrightLicensedContent",
			"PageViewNoteHateHarassment",
			"PageViewNoteViolenceGore",
			"PageViewNoteSelfHarmSuicide",
			"PageViewNotePrivacyViolationPersonalInfo",
			"PageViewNoteIllegalActivityDrugs",
			"PageViewNoteOther",
			"PageViewNoteEmail",
			"PageViewNoteDetails",
			"PageViewNoteDescribeTheIssue",
			"PageViewNoteSend",
		},
		"edit_note.go.html": {
			"PageEditNoteDescription",
			"PageEditNoteKeywords",
			"PageEditNoteOGDescription",
			"PageEditNoteTitle",
			"PageEditNoteYourName",
			"PageEditNoteYourNote",
			"PageEditNoteMenuPH",
			"PageEditNotePublish",
			"PageEditNotePublishPassword",
			"PageEditNoteModalTitle",
			"PageEditNoteModalMessage",
			"PageEditNoteModalYourPass",
			"MsgErrEmptyTitle",
			"MsgErrTitleSmall",
			"MsgErrContentEmpty",
			"MsgErrContentLarge",
			"MsgErrMenuLarge",
		},
	}
)

func AddMessages(c *gin.Context, name string, params ...gin.H) gin.H {
	lang, ok := c.Get("lang")
	if !ok {
		lang = "en"
	}

	h := gin.H{
		"Lang": lang,
	}

	for _, msg := range pageMsgs[name] {
		h[msg] = template.HTML(T(c, msg))
	}

	for _, param := range params {
		for key, arg := range param {
			h[key] = arg
		}
	}

	return h
}
