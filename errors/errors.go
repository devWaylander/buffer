package errors

const (
	ErrUnauthorized = "ERR_INVALID_TOKEN"
	ErrDecodeReq    = "ERR_FAILED_TO_DECODE_JSON_REQ"
	ErrMarshalResp  = "ERR_FAILED_TO_ENCODE_JSON_RESP"

	ErrAuthUserIDIsMissing        = "ERR_AUTH_USER_ID_IS_MISSING"
	ErrCommentIsMissing           = "ERR_COMMENT_IS_MISSING"
	ErrFactTimeIsMissing          = "ERR_FACT_TIME_IS_MISSING"
	ErrIndicatorToFactIDIsMissing = "ERR_INDICATOR_TO_FACT_ID_IS_MISSING"
	ErrIndicatorToMoIDIsMissing   = "ERR_INDICATOR_TO_MO_ID_IS_MISSING"
	ErrIsPlanIsMissing            = "ERR_IS_PLAN_IS_MISSING"
	ErrPeriodEndIsMissing         = "ERR_PERIOD_END_IS_MISSING"
	ErrPeriodKeyIsMissing         = "ERR_PERIOD_KEY_IS_MISSING"
	ErrPeriodStartIsMissing       = "ERR_PERIOD_START_IS_MISSING"
	ErrValueIsMissing             = "ERR_VALUE_IS_MISSING"
)
