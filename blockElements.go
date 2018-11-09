package slack

/*
TODO:
[ ] enforce limit on Overflow 10 options limit
[ ] enforce limit on Select 10 options limit

*/

// BlockElementType type of elements that could be contained in a Block
type BlockElementType string

const (
	// ButtonBlockElementType a button element
	ButtonBlockElementType BlockElementType = "button"
	// DatePickerBlockElementType a datePicker element
	DatePickerBlockElementType BlockElementType = "datepicker"
	// OverflowBlockElementType an overflow menu
	OverflowBlockElementType BlockElementType = "overflow"
	// ImageBlockElementType an image block element
	ImageBlockElementType BlockElementType = "image"
	//PlainTextBlockElementType to render a plaintext string
	PlainTextBlockElementType BlockElementType = "plaintext"
	//MrkdwnBlockElementType type to hold a markDown string
	MrkdwnBlockElementType BlockElementType = "mrkdwn"
	// UserBlockElementType a user block
	UserBlockElementType BlockElementType = "user"

	// ** Select types **

	// SelectBlockElementType a select element
	SelectBlockElementType BlockElementType = "select"
	// SelectStaticBlockElementType for a static list type menu
	SelectStaticBlockElementType BlockElementType = "static_select"
	// SelectExternalBlockElementType for a dynamic list type menu
	SelectExternalBlockElementType BlockElementType = "external_select"
	// SelectUsersBlockElementType for a users list type menu
	SelectUsersBlockElementType BlockElementType = "users_select"
	// SelectConversationsBlockElementType for a conversations list type menu
	SelectConversationsBlockElementType BlockElementType = "conversations_select"
	// SelectChannelsBlockElementType for a channels list type menu
	SelectChannelsBlockElementType BlockElementType = "channels_select"
)

// BlockElement generic BlockElement type
type BlockElement interface{}

// TextBlockElement to wrap `PlainTextBlockElement` and `MrkdwnBlockElement`
type TextBlockElement interface {
	textContent() string
}

// PlainTextBlockElement basic plain text with a flag to render emoji
type PlainTextBlockElement struct {
	Type  BlockElementType `json:"type"`
	Text  string           `json:"text"`
	Emoji bool             `json:"emoji"`
}

// NewPlainText creates a new PlainTextBlockElement with a Emoji=true
func NewPlainText(text string) PlainTextBlockElement {
	return NewPlainTextBlockElement(text, true)
}

// NewPlainTextBlockElement creates a new plainText element
func NewPlainTextBlockElement(text string, renderEmoji bool) PlainTextBlockElement {
	return PlainTextBlockElement{
		Type:  PlainTextBlockElementType,
		Text:  text,
		Emoji: renderEmoji,
	}
}

func (p PlainTextBlockElement) textContent() string {
	return p.Text
}

// MrkdwnBlockElement Markdown text element
type MrkdwnBlockElement struct {
	Type  BlockElementType `json:"type"`
	Text  string           `json:"text"`
	Parse string           `json:"parse"`
}

// NewMrkdwnBlockElement creates a new plainText element
func NewMrkdwnBlockElement(text string, parse bool) MrkdwnBlockElement {
	shouldParse := "none"
	if parse {
		shouldParse = "full"
	}
	return MrkdwnBlockElement{
		Type:  MrkdwnBlockElementType,
		Text:  text,
		Parse: shouldParse,
	}
}

func (m MrkdwnBlockElement) textContent() string {
	return m.Text
}

// SectionBlockCompatibleElement an interface to check if a BlockElement can be embedded in `TextBlock`
type SectionBlockCompatibleElement interface {
	canEmbeddInSectionBlock() bool
}

// ButtonBlockElement a button action block element
type ButtonBlockElement struct {
	Type     BlockElementType      `json:"type"`
	Text     PlainTextBlockElement `json:"text"`
	ActionID string                `json:"action_id,omitempty"`
	Value    string                `json:"value,omitempty"`
	URL      string                `json:"url,omitempty"`
}

func (b ButtonBlockElement) canEmbeddInSectionBlock() bool {
	return true
}

// NewButtonBlockElement convnience method to create `ButtonBlockElement`
func NewButtonBlockElement(text, actionID, value string) ButtonBlockElement {
	plainText := NewPlainText(text)
	return ButtonBlockElement{
		Type:     ButtonBlockElementType,
		Text:     plainText,
		ActionID: actionID,
		Value:    value,
	}
}

// OverflowBlockElement a small context menu of up to 10 options
type OverflowBlockElement struct {
	Type     BlockElementType    `json:"type"`
	ActionID string              `json:"action_id"`
	Options  []BlockSelectOption `json:"options"` //Limited to 10
}

func (o OverflowBlockElement) canEmbeddInSectionBlock() bool {
	return true
}

// NewOverflowBlockElement convnience method to create `OverflowBlockElement`
func NewOverflowBlockElement(actionID string, options []BlockSelectOption) OverflowBlockElement {
	return OverflowBlockElement{
		Type:     OverflowBlockElementType,
		ActionID: actionID,
		Options:  options,
	}
}

// SelectBlockDataSource types of select datasource
type SelectBlockDataSource string

const (
	// StaticSelectDataSource menu with static Options/OptionGroups
	StaticSelectDataSource SelectBlockDataSource = "static"
	// ExternalSelectDataSource dynamic datasource
	ExternalSelectDataSource SelectBlockDataSource = "external"
	// ConversationsSelectDataSource provides a list of conversations
	ConversationsSelectDataSource SelectBlockDataSource = "conversations"
	// ChannelsSelectDataSource provides a list of channels
	ChannelsSelectDataSource SelectBlockDataSource = "channels"
	// UsersSelectDataSource provides a list of users
	UsersSelectDataSource SelectBlockDataSource = "users"
	// DateSelectDataSource a date picker data source
	DateSelectDataSource SelectBlockDataSource = "date"
)

// SelectBlockElement is a `Select Element` of type `select`
type SelectBlockElement struct {
	Type        BlockElementType      `json:"type"`
	ActionID    string                `json:"action_id,omitempty"`
	PlaceHolder PlainTextBlockElement `json:"placeholder,omitempty"`
	DataSource  SelectBlockDataSource `json:"data_source,omitempty"`
	Options     []BlockSelectOption   `json:"options,omitempty"`
}

func (s SelectBlockElement) canEmbeddInSectionBlock() bool {
	return true
}

// NewSelectBlockElement convenience Constructor to create a `SelectBlockElement`
func NewSelectBlockElement(actionID, placeHolder string, dataSource SelectBlockDataSource, options []BlockSelectOption) SelectBlockElement {
	placeHolderPlainText := NewPlainText(placeHolder)
	return SelectBlockElement{
		Type:        SelectBlockElementType,
		ActionID:    actionID,
		PlaceHolder: placeHolderPlainText,
		DataSource:  dataSource,
		Options:     options,
	}
}

// ImageBlockElement an image block
type ImageBlockElement struct {
	Type     BlockElementType `json:"type"`
	ImageURL string           `json:"image_url,omitempty"`
	AltText  string           `json:"alt_text,omitempty"`
}

func (i ImageBlockElement) canEmbeddInSectionBlock() bool {
	return true
}

// NewImageBlockElemenet convnience method to create `ImageBlockElement`
func NewImageBlockElemenet(imageURL, altText string) ImageBlockElement {
	return ImageBlockElement{
		Type:     ImageBlockElementType,
		ImageURL: imageURL,
		AltText:  altText,
	}
}

// DatePickerBlockElement a datepicker element type
type DatePickerBlockElement struct {
	Type        BlockElementType         `json:"type"`
	ActionID    string                   `json:"action_id,omitempty"`
	Value       string                   `json:"value,omitempty"`
	InitialDate string                   `json:"initial_date,omitempty"`
	PlaceHolder *PlainTextBlockElement   `json:"placeholder,omitempty"`
	Confirm     *BlockActionConfirmation `json:"confirm,omitempty"`
}

func (d DatePickerBlockElement) canEmbeddInSectionBlock() bool {
	return true
}

// NewDatePickerBlockElement convnience method to create `DatePickerBlockElement`
func NewDatePickerBlockElement(actionID, value string) DatePickerBlockElement {
	return DatePickerBlockElement{
		Type:     DatePickerBlockElementType,
		ActionID: actionID,
		Value:    value,
	}
}

// BlockActionConfirmation will popup a confirmation before an action
type BlockActionConfirmation struct {
	Text    TextBlockElement      `json:"text"`
	Title   PlainTextBlockElement `json:"title,omitempty"`
	Confirm PlainTextBlockElement `json:"confirm,omitempty"`
	Deny    PlainTextBlockElement `json:"deny,omitempty"`
}

// UserBlockElement a `User` block element
type UserBlockElement struct {
	Type   BlockElementType `json:"type"`
	UserID string           `json:"user_id"`
}

// NewUserBlockElement creates a new `UserBlockElement`
func NewUserBlockElement(userID string) UserBlockElement {
	return UserBlockElement{
		Type:   UserBlockElementType,
		UserID: userID,
	}
}
