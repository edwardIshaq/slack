package slack

// BlockType captures the type of block
type BlockType string

const (
	// DividerBlockType horizontal divider
	DividerBlockType BlockType = "divider"
	// ImageBlockType an image block
	ImageBlockType BlockType = "image"
	// SectionBlockType a simple text block
	SectionBlockType BlockType = "section"
	// ActionsBlockType is a collection of up to 5 Actions (select, button, overflow, datepicker)
	ActionsBlockType BlockType = "actions"
	// ContextBlockType is a collection of up to 10 elements of type (image, text, user)
	ContextBlockType BlockType = "context"
)

// Block interface for the attachments
type Block interface {
	GetBlockID() string
	GetType() BlockType
}

// BaseBlock base block element
type BaseBlock struct {
	BlockID string    `json:"block_id,omitempty"`
	Type    BlockType `json:"type"`
}

// GetBlockID returns the block's ID
func (b BaseBlock) GetBlockID() string {
	return b.BlockID
}

// GetType returns the block's Type
func (b BaseBlock) GetType() BlockType {
	return b.Type
}

// DividerBlock a horizontal devider block
type DividerBlock struct {
	BaseBlock
}

// NewDividerBlock builds a devider block
func NewDividerBlock(blockID string) DividerBlock {
	return DividerBlock{
		BaseBlock: BaseBlock{
			BlockID: blockID,
			Type:    DividerBlockType,
		},
	}
}

// SectionBlock a simple text block
type SectionBlock struct {
	BaseBlock
	Text    TextBlockElement              `json:"text,omitempty"`
	Element SectionBlockCompatibleElement `json:"accessory,omitempty"`
}

// NewSectionBlockWithText creates a new SectionBlock with a PlainText
func NewSectionBlockWithText(text string) SectionBlock {
	plainText := NewPlainTextBlockElement(text, true)
	return SectionBlock{
		BaseBlock: BaseBlock{
			Type: SectionBlockType,
		},
		Text: plainText,
	}
}

// NewSectionBlock constructs a new text block with ID and text
func NewSectionBlock(textElement TextBlockElement, blockID string) SectionBlock {
	return SectionBlock{
		BaseBlock: BaseBlock{
			BlockID: blockID,
			Type:    SectionBlockType,
		},
		Text: textElement,
	}
}

// NewSectionBlockWithElement create new SectionBlock with Embedded element
func NewSectionBlockWithElement(text TextBlockElement, blockID string, element SectionBlockCompatibleElement) SectionBlock {
	return SectionBlock{
		BaseBlock: BaseBlock{
			BlockID: blockID,
			Type:    SectionBlockType,
		},
		Text:    text,
		Element: element,
	}
}

// ActionBlock base block element
type ActionBlock struct {
	BaseBlock
	Elements []BlockElement `json:"elements"`
}

// NewActionBlock constructor
func NewActionBlock(blockID string, elements []BlockElement) ActionBlock {
	return ActionBlock{
		BaseBlock: BaseBlock{
			BlockID: blockID,
			Type:    ActionsBlockType,
		},
		Elements: elements,
	}
}

/*
{
	"type": "image",
	"block_id": "image4",
	"image_url": "https://scontent-sjc3-1.cdninstagram.com/vp/64d7aa4ab1a55892036c52b2237f3868/5B97948F/t51.2885-15/s640x640/sh0.08/e35/c0.135.1080.1080/26155058_127776921353331_485398838513762304_n.jpg",
	"alt_text": "cat",
	"caption": "Bubsy eats a cheeto"
}
*/

// ImageBlock a text block with an image embedded in it
type ImageBlock struct {
	BaseBlock
	ImageURL string                `json:"image_url,omitempty"`
	Title    PlainTextBlockElement `json:"title,omitempty"`
	AltText  string                `json:"alt_text,omitempty"`
}

// NewImageBlock constructs a new text block with ID and text
func NewImageBlock(imageURL, altText, blockID string, title PlainTextBlockElement) ImageBlock {
	return ImageBlock{
		BaseBlock: BaseBlock{
			BlockID: blockID,
			Type:    ImageBlockType,
		},
		ImageURL: imageURL,
		AltText:  altText,
		Title:    title,
	}
}

// ContextBlock provides a context for an image/plainText/mrkdwn/user
type ContextBlock struct {
	BaseBlock
	Elements []BlockElement `json:"elements"` //max 10
}

// NewContextBlock initializes a new `ContextBlock`
func NewContextBlock(elements []BlockElement) ContextBlock {
	return ContextBlock{
		BaseBlock: BaseBlock{
			Type: ContextBlockType,
		},
		Elements: elements,
	}
}
