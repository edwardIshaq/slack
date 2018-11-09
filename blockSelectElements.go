package slack

// BlockSelectOption the individual option to appear in a `Select` Block element
type BlockSelectOption struct {
	Text        PlainTextBlockElement `json:"text"` // Required.
	Value       string                `json:"value,omitempty"`
	URL         string                `json:"url,omitempty"`
	Description string                `json:"description,omitempty"`
}

// BlockOptionGroup is a group of `BlockSelectOption`
type BlockOptionGroup struct {
	Label   PlainTextBlockElement `json:"label"`
	Options []BlockSelectOption   `json:"options"`
}

//-----------------------------------------------------------------------------------
// ----------------------------- BaseSelectBlockElement -----------------------------
//-----------------------------------------------------------------------------------

// BaseSelectBlockElement base object for all common select types
type BaseSelectBlockElement struct {
	Type        BlockElementType         `json:"type"`
	ActionID    string                   `json:"action_id,omitempty"`
	Placeholder *PlainTextBlockElement   `json:"placeholder,omitempty"`
	Confirm     *BlockActionConfirmation `json:"confirm,omitempty"`
}

func (c BaseSelectBlockElement) canEmbeddInSectionBlock() bool {
	return true
}

//-----------------------------------------------------------------------------------
// ----------------------------- StaticSelectBlockElement ---------------------------
//-----------------------------------------------------------------------------------

// StaticSelectBlockElement provides its own static options/options_groups
type StaticSelectBlockElement struct {
	BaseSelectBlockElement
	InitialOption *BlockSelectOption  `json:"initial_option,omitempty"`
	Options       []BlockSelectOption `json:"options,omitempty"`
	OptionGroups  []BlockOptionGroup  `json:"option_groups,omitempty"`
}

// NewStaticSelectBlockElement prepares a new `StaticSelectBlockElement`
func NewStaticSelectBlockElement() StaticSelectBlockElement {
	return StaticSelectBlockElement{
		BaseSelectBlockElement: BaseSelectBlockElement{
			Type: SelectStaticBlockElementType,
		},
	}
}

//-----------------------------------------------------------------------------------
// ----------------------------- ExternalSelectBlockElement ---------------------------
//-----------------------------------------------------------------------------------

// ExternalSelectBlockElement a select element with a dynamic remote dataSource
type ExternalSelectBlockElement struct {
	BaseSelectBlockElement
	InitialOption  *BlockSelectOption `json:"initial_option,omitempty"`
	MinQueryLength int                `json:"min_query_length,omitempty"`
}

// NewExternalSelectBlockElement prepares a new `ExternalSelectBlockElement`
func NewExternalSelectBlockElement() ExternalSelectBlockElement {
	return ExternalSelectBlockElement{
		BaseSelectBlockElement: BaseSelectBlockElement{
			Type: SelectExternalBlockElementType,
		},
	}
}

//-----------------------------------------------------------------------------------
// ----------------------------- ConversationsSelectBlockElement ---------------------------
//-----------------------------------------------------------------------------------

// ConversationsSelectBlockElement a conversations select element
type ConversationsSelectBlockElement struct {
	BaseSelectBlockElement
	InitialOption  *BlockSelectOption `json:"initial_option,omitempty"`
	MinQueryLength int                `json:"min_query_length,omitempty"`
}

// MakeConversationsSelectBlockElement prepares a new `ConversationsSelectBlockElement`
func MakeConversationsSelectBlockElement() ConversationsSelectBlockElement {
	return ConversationsSelectBlockElement{
		BaseSelectBlockElement: BaseSelectBlockElement{
			Type: SelectConversationsBlockElementType,
		},
	}
}

// NewConversationsSelectBlockElement creates a new Conversations Select block element
func NewConversationsSelectBlockElement(actionID, initialConersation, placeholder string) ConversationsSelectBlockElement {
	placeholderPlaintext := NewPlainText(placeholder)
	return ConversationsSelectBlockElement{
		BaseSelectBlockElement: BaseSelectBlockElement{
			Type:        SelectConversationsBlockElementType,
			ActionID:    actionID,
			Placeholder: &placeholderPlaintext,
		},
	}
}

//-----------------------------------------------------------------------------------
// ----------------------------- ChannelsSelectBlockElement ---------------------------
//-----------------------------------------------------------------------------------

// ChannelsSelectBlockElement a conversations select element
type ChannelsSelectBlockElement struct {
	BaseSelectBlockElement
	InitialChannelID string `json:"initial_channel,omitempty"`
}

// MakeChannelsSelectBlockElement prepares a `ChannelsSelectBlockElement`
func MakeChannelsSelectBlockElement() ChannelsSelectBlockElement {
	return ChannelsSelectBlockElement{
		BaseSelectBlockElement: BaseSelectBlockElement{
			Type: SelectChannelsBlockElementType,
		},
	}
}

//-----------------------------------------------------------------------------------
// ----------------------------- UsersSelectBlockElement ---------------------------
//-----------------------------------------------------------------------------------

// UsersSelectBlockElement a conversations select element
type UsersSelectBlockElement struct {
	BaseSelectBlockElement
	InitialUserID string `json:"initial_user,omitempty"`
}

// MakeUsersSelectBlockElement prepares a `UsersSelectBlockElement`
func MakeUsersSelectBlockElement() UsersSelectBlockElement {
	return UsersSelectBlockElement{
		BaseSelectBlockElement: BaseSelectBlockElement{
			Type: SelectUsersBlockElementType,
		},
	}
}
