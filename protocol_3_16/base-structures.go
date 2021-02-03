package protocol

import (
	"encoding/json"
)

// https://microsoft.github.io/language-server-protocol/specification#uri

type DocumentUri = string

type URI = string

// https://microsoft.github.io/language-server-protocol/specification#regExp

/**
 * Client capabilities specific to regular expressions.
 */
type RegularExpressionsClientCapabilities struct {
	/**
	 * The engine's name.
	 */
	Engine string `json:"engine"`

	/**
	 * The engine's version.
	 */
	Version *string `json:"version,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specification#textDocuments

var EOL = []string{"\n", "\r\n", "\r"}

// https://microsoft.github.io/language-server-protocol/specification#position

type Position struct {
	/**
	 * Line position in a document (zero-based).
	 */
	Line UInteger `json:"line"`

	/**
	 * Character offset on a line in a document (zero-based). Assuming that
	 * the line is represented as a string, the `character` value represents
	 * the gap between the `character` and `character + 1`.
	 *
	 * If the character value is greater than the line length it defaults back
	 * to the line length.
	 */
	Character UInteger `json:"character"`
}

// https://microsoft.github.io/language-server-protocol/specification#range

type Range struct {
	/**
	 * The range's start position.
	 */
	Start Position `json:"start"`

	/**
	 * The range's end position.
	 */
	End Position `json:"end"`
}

// https://microsoft.github.io/language-server-protocol/specification#location

type Location struct {
	URI   DocumentUri `json:"uri"`
	Range Range       `json:"range"`
}

// https://microsoft.github.io/language-server-protocol/specification#locationLink

type LocationLink struct {
	/**
	 * Span of the origin of this link.
	 *
	 * Used as the underlined span for mouse interaction. Defaults to the word
	 * range at the mouse position.
	 */
	OriginSelectionRange *Range `json:"uri,omitempty"`

	/**
	 * The target resource identifier of this link.
	 */
	TargetURI DocumentUri `json:"targetUri"`

	/**
	 * The full target range of this link. If the target for example is a symbol
	 * then target range is the range enclosing this symbol not including
	 * leading/trailing whitespace but everything else like comments. This
	 * information is typically used to highlight the range in the editor.
	 */
	TargetRange Range `json:"targetRange"`

	/**
	 * The range that should be selected and revealed when this link is being
	 * followed, e.g the name of a function. Must be contained by the the
	 * `targetRange`. See also `DocumentSymbol#range`
	 */
	TargetSelectionRange Range `json:"targetSelectionRage"`
}

// https://microsoft.github.io/language-server-protocol/specification#diagnostic

type Diagnostic struct {
	/**
	 * The range at which the message applies.
	 */
	Range Range `json:"range"`

	/**
	 * The diagnostic's severity. Can be omitted. If omitted it is up to the
	 * client to interpret diagnostics as error, warning, info or hint.
	 */
	Severity *DiagnosticSeverity `json:"severity,omitempty"`

	/**
	 * The diagnostic's code, which might appear in the user interface.
	 */
	Code *IntegerOrString `json:"code,omitempty"`

	/**
	 * An optional property to describe the error code.
	 *
	 * @since 3.16.0
	 */
	CodeDescription *CodeDescription `json:"codeDescription,omitempty"`

	/**
	 * A human-readable string describing the source of this
	 * diagnostic, e.g. 'typescript' or 'super lint'.
	 */
	Source *string `json:"source,omitempty"`

	/**
	 * The diagnostic's message.
	 */
	Message string `json:"message"`

	/**
	 * Additional metadata about the diagnostic.
	 *
	 * @since 3.15.0
	 */
	Tags []DiagnosticTag `json:"tags,omitempty"`

	/**
	 * An array of related diagnostic information, e.g. when symbol-names within
	 * a scope collide all definitions can be marked via this property.
	 */
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`

	/**
	 * A data entry field that is preserved between a
	 * `textDocument/publishDiagnostics` notification and
	 * `textDocument/codeAction` request.
	 *
	 * @since 3.16.0
	 */
	Data interface{} `json:"data,omitempty"`
}

type DiagnosticSeverity Integer

const (
	/**
	 * Reports an error.
	 */
	DiagnosticSeverityError = DiagnosticSeverity(1)

	/**
	 * Reports a warning.
	 */
	DiagnosticSeverityWarning = DiagnosticSeverity(2)

	/**
	 * Reports an information.
	 */
	DiagnosticSeverityInformation = DiagnosticSeverity(3)

	/**
	 * Reports a hint.
	 */
	DiagnosticSeverityHint = DiagnosticSeverity(4)
)

/**
 * The diagnostic tags.
 *
 * @since 3.15.0
 */
type DiagnosticTag Integer

const (
	/**
	 * Unused or unnecessary code.
	 *
	 * Clients are allowed to render diagnostics with this tag faded out
	 * instead of having an error squiggle.
	 */
	DiagnosticTagUnnecessary = DiagnosticTag(1)

	/**
	 * Deprecated or obsolete code.
	 *
	 * Clients are allowed to rendered diagnostics with this tag strike through.
	 */
	DiagnosticTagDeprecated = DiagnosticTag(2)
)

/**
 * Represents a related message and source code location for a diagnostic.
 * This should be used to point to code locations that cause or are related to
 * a diagnostics, e.g when duplicating a symbol in a scope.
 */
type DiagnosticRelatedInformation struct {
	/**
	 * The location of this related diagnostic information.
	 */
	Location Location `json:"location"`

	/**
	 * The message of this related diagnostic information.
	 */
	Message string `json:"message"`
}

/**
 * Structure to capture a description for an error code.
 *
 * @since 3.16.0
 */
type CodeDescription struct {
	/**
	 * An URI to open with more information about the diagnostic error.
	 */
	HRef URI `json:"href"`
}

// https://microsoft.github.io/language-server-protocol/specification#command

type Command struct {
	/**
	 * Title of the command, like `save`.
	 */
	Title string `json:"title"`

	/**
	 * The identifier of the actual command handler.
	 */
	Command string `json:"command"`

	/**
	 * Arguments that the command handler should be
	 * invoked with.
	 */
	Arguments []interface{} `json:"arguments,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specification#textEdit

type TextEdit struct {
	/**
	 * The range of the text document to be manipulated. To insert
	 * text into a document create a range where start === end.
	 */
	Range Range `json:"range"`

	/**
	 * The string to be inserted. For delete operations use an
	 * empty string.
	 */
	NewText string `json:"newText"`
}

/**
 * Additional information that describes document changes.
 *
 * @since 3.16.0
 */
type ChangeAnnotation struct {
	/**
	 * A human-readable string describing the actual change. The string
	 * is rendered prominent in the user interface.
	 */
	Label string `json:"label"`

	/**
	 * A flag which indicates that user confirmation is needed
	 * before applying the change.
	 */
	NeedsConfirmation *bool `json:"needsConfirmation,omitempty"`

	/**
	 * A human-readable string which is rendered less prominent in
	 * the user interface.
	 */
	Description *string `json:"description,omitempty"`
}

/**
 * An identifier referring to a change annotation managed by a workspace
 * edit.
 *
 * @since 3.16.0
 */
type ChangeAnnotationIdentifier = string

/**
 * A special text edit with an additional change annotation.
 *
 * @since 3.16.0
 */
type AnnotatedTextEdit struct {
	TextEdit

	/**
	 * The actual annotation identifier.
	 */
	AnnotationID ChangeAnnotationIdentifier `json:"annotationId"`
}

// https://microsoft.github.io/language-server-protocol/specification#textDocumentEdit

type TextDocumentEdit struct {
	/**
	 * The text document to change.
	 */
	TextDocument OptionalVersionedTextDocumentIdentifier `json:"textDocument"`

	/**
	 * The edits to be applied.
	 *
	 * @since 3.16.0 - support for AnnotatedTextEdit. This is guarded by the
	 * client capability `workspace.workspaceEdit.changeAnnotationSupport`
	 */
	Edits []interface{} `json:"edits"` // TextEdit | AnnotatedTextEdit
}

// json.Unmarshaler interface
func (self *TextDocumentEdit) UnmarshalJSON(data []byte) error {
	var value struct {
		TextDocument OptionalVersionedTextDocumentIdentifier `json:"textDocument"`
		Edits        []json.RawMessage                       `json:"edits"` // TextEdit | AnnotatedTextEdit
	}

	if err := json.Unmarshal(data, &value); err == nil {
		self.TextDocument = value.TextDocument

		for _, edit := range value.Edits {
			var value TextEdit
			if err = json.Unmarshal(edit, &value); err == nil {
				self.Edits = append(self.Edits, value)
			} else {
				var value AnnotatedTextEdit
				if err = json.Unmarshal(edit, &value); err == nil {
					self.Edits = append(self.Edits, value)
				} else {
					return err
				}
			}
		}

		return nil
	} else {
		return err
	}
}

// https://microsoft.github.io/language-server-protocol/specification#resourceChanges

/**
 * Options to create a file.
 */
type CreateFileOptions struct {
	/**
	 * Overwrite existing file. Overwrite wins over `ignoreIfExists`
	 */
	Overwrite *bool `json:"overwrite,omitempty"`

	/**
	 * Ignore if exists.
	 */
	IgnoreIfExists *bool `json:"ignoreIfExists,omitempty"`
}

/**
 * Create file operation
 */
type CreateFile struct {
	/**
	 * A create
	 */
	Kind string `json:"kind"` // == "create"

	/**
	 * The resource to create.
	 */
	URI DocumentUri `json:"uri"`

	/**
	 * Additional options
	 */
	Options *CreateFileOptions `json:"options,omitempty"`

	/**
	 * An optional annotation identifer describing the operation.
	 *
	 * @since 3.16.0
	 */
	AnnotationID *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
}

/**
 * Rename file options
 */
type RenameFileOptions struct {
	/**
	 * Overwrite target if existing. Overwrite wins over `ignoreIfExists`
	 */
	Overwrite *bool `json:"overwrite,omitempty"`

	/**
	 * Ignores if target exists.
	 */
	IgnoreIfExists *bool `json:"ignoreIfExists,omitempty"`
}

/**
 * Rename file operation
 */
type RenameFile struct {
	/**
	 * A rename
	 */
	Kind string `json:"kind"` // == "rename"

	/**
	 * The old (existing) location.
	 */
	OldURI DocumentUri `json:"oldUri"`

	/**
	 * The new location.
	 */
	NewURI DocumentUri `json:"newUri"`

	/**
	 * Rename options.
	 */
	Options *RenameFileOptions `json:"options,omitempty"`

	/**
	 * An optional annotation identifer describing the operation.
	 *
	 * @since 3.16.0
	 */
	AnnotationID *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
}

/**
 * Delete file options
 */
type DeleteFileOptions struct {
	/**
	 * Delete the content recursively if a folder is denoted.
	 */
	Recursive *bool `json:"recursive,omitempty"`

	/**
	 * Ignore the operation if the file doesn't exist.
	 */
	IgnoreIfNotExists *bool `json:"ignoreIfNotExists,omitempty"`
}

/**
 * Delete file operation
 */
type DeleteFile struct {
	/**
	 * A delete
	 */
	Kind string `json:"kind"` // == "delete"

	/**
	 * The file to delete.
	 */
	URI DocumentUri `json:"uri"`

	/**
	 * Delete options.
	 */
	Options *DeleteFileOptions `json:"options,omitempty"`

	/**
	 * An optional annotation identifer describing the operation.
	 *
	 * @since 3.16.0
	 */
	AnnotationID *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specification#workspaceEdit

type WorkspaceEdit struct {
	/**
	 * Holds changes to existing resources.
	 */
	Changes map[DocumentUri][]TextEdit `json:"changes,omitempty"`

	/**
	 * Depending on the client capability
	 * `workspace.workspaceEdit.resourceOperations` document changes are either
	 * an array of `TextDocumentEdit`s to express changes to n different text
	 * documents where each text document edit addresses a specific version of
	 * a text document. Or it can contain above `TextDocumentEdit`s mixed with
	 * create, rename and delete file / folder operations.
	 *
	 * Whether a client supports versioned document edits is expressed via
	 * `workspace.workspaceEdit.documentChanges` client capability.
	 *
	 * If a client neither supports `documentChanges` nor
	 * `workspace.workspaceEdit.resourceOperations` then only plain `TextEdit`s
	 * using the `changes` property are supported.
	 */
	DocumentChanges []interface{} `json:"documentChanges,omitempty"` // TextDocumentEdit | CreateFile | RenameFile | DeleteFile

	/**
	 * A map of change annotations that can be referenced in
	 * `AnnotatedTextEdit`s or create, rename and delete file / folder
	 * operations.
	 *
	 * Whether clients honor this property depends on the client capability
	 * `workspace.changeAnnotationSupport`.
	 *
	 * @since 3.16.0
	 */
	ChangeAnnotations map[ChangeAnnotationIdentifier]ChangeAnnotation `json:"changeAnnotations,omitempty"`
}

// json.Unmarshaler interface
func (self *WorkspaceEdit) UnmarshalJSON(data []byte) error {
	var value struct {
		Changes           map[DocumentUri][]TextEdit                      `json:"changes"`
		DocumentChanges   []json.RawMessage                               `json:"documentChanges"` // TextDocumentEdit | CreateFile | RenameFile | DeleteFile
		ChangeAnnotations map[ChangeAnnotationIdentifier]ChangeAnnotation `json:"changeAnnotations"`
	}

	if err := json.Unmarshal(data, &value); err == nil {
		self.Changes = value.Changes
		self.ChangeAnnotations = value.ChangeAnnotations

		for _, documentChange := range value.DocumentChanges {
			var value TextDocumentEdit
			if err = json.Unmarshal(documentChange, &value); err == nil {
				self.DocumentChanges = append(self.DocumentChanges, value)
			} else {
				var value CreateFile
				if err = json.Unmarshal(documentChange, &value); err == nil {
					self.DocumentChanges = append(self.DocumentChanges, value)
				} else {
					var value RenameFile
					if err = json.Unmarshal(documentChange, &value); err == nil {
						self.DocumentChanges = append(self.DocumentChanges, value)
					} else {
						var value DeleteFile
						if err = json.Unmarshal(documentChange, &value); err == nil {
							self.DocumentChanges = append(self.DocumentChanges, value)
						} else {
							return err
						}
					}
				}
			}
		}

		return nil
	} else {
		return err
	}
}

// https://microsoft.github.io/language-server-protocol/specification#workspaceEditClientCapabilities

type WorkspaceEditClientCapabilities struct {
	/**
	 * The client supports versioned document changes in `WorkspaceEdit`s
	 */
	DocumentChanges *bool `json:"documentChanges,omitempty"`

	/**
	 * The resource operations the client supports. Clients should at least
	 * support 'create', 'rename' and 'delete' files and folders.
	 *
	 * @since 3.13.0
	 */
	ResourceOperations []ResourceOperationKind `json:"resourceOperations,omitempty"`

	/**
	 * The failure handling strategy of a client if applying the workspace edit
	 * fails.
	 *
	 * @since 3.13.0
	 */
	FailureHandling *FailureHandlingKind `json:"failureHandling,omitempty"`

	/**
	 * Whether the client normalizes line endings to the client specific
	 * setting.
	 * If set to `true` the client will normalize line ending characters
	 * in a workspace edit to the client specific new line character(s).
	 *
	 * @since 3.16.0
	 */
	NormalizesLineEndings *bool `json:"normalizesLineEndings,omitempty"`

	/**
	 * Whether the client in general supports change annotations on text edits,
	 * create file, rename file and delete file changes.
	 *
	 * @since 3.16.0
	 */
	ChangeAnnotationSupport struct {
		/**
		 * Whether the client groups edits with equal labels into tree nodes,
		 * for instance all edits labelled with "Changes in Strings" would
		 * be a tree node.
		 */
		GroupsOnLabel *bool `json:"groupsOnLabel,omitempty"`
	} `json:"changeAnnotationSupport,omitempty"`
}

/**
 * The kind of resource operations supported by the client.
 */
type ResourceOperationKind string

const (
	/**
	 * Supports creating new files and folders.
	 */
	ResourceOperationKindCreate = ResourceOperationKind("create")

	/**
	 * Supports renaming existing files and folders.
	 */
	ResourceOperationKindRename = ResourceOperationKind("rename")

	/**
	 * Supports deleting existing files and folders.
	 */
	ResourceOperationKindDelete = ResourceOperationKind("delete")
)

type FailureHandlingKind string

const (
	/**
	 * Applying the workspace change is simply aborted if one of the changes
	 * provided fails. All operations executed before the failing operation
	 * stay executed.
	 */
	FailureHandlingKindAbort = FailureHandlingKind("abort")

	/**
	 * All operations are executed transactional. That means they either all
	 * succeed or no changes at all are applied to the workspace.
	 */
	FailureHandlingKindTransactional = FailureHandlingKind("transactional")

	/**
	 * If the workspace edit contains only textual file changes they are
	 * executed transactional. If resource changes (create, rename or delete
	 * file) are part of the change the failure handling strategy is abort.
	 */
	FailureHandlingKindTextOnlyTransactional = FailureHandlingKind("textOnlyTransactional")

	/**
	 * The client tries to undo the operations already executed. But there is no
	 * guarantee that this is succeeding.
	 */
	FailureHandlingKindUndo = FailureHandlingKind("undo")
)

// https://microsoft.github.io/language-server-protocol/specification#textDocumentIdentifier

type TextDocumentIdentifier struct {
	/**
	 * The text document's URI.
	 */
	URI DocumentUri `json:"uri"`
}

// https://microsoft.github.io/language-server-protocol/specification#textDocumentItem

type TextDocumentItem struct {
	/**
	 * The text document's URI.
	 */
	URI DocumentUri `json:"uri"`

	/**
	 * The text document's language identifier.
	 */
	LanguageID string `json:"languageId"`

	/**
	 * The version number of this document (it will increase after each
	 * change, including undo/redo).
	 */
	Version Integer `json:"version"`

	/**
	 * The content of the opened text document.
	 */
	Text string `json:"text"`
}

// https://microsoft.github.io/language-server-protocol/specification#versionedTextDocumentIdentifier

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier

	/**
	 * The version number of this document.
	 *
	 * The version number of a document will increase after each change,
	 * including undo/redo. The number doesn't need to be consecutive.
	 */
	Version Integer `json:"version"`
}

type OptionalVersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier

	/**
	 * The version number of this document. If an optional versioned text document
	 * identifier is sent from the server to the client and the file is not
	 * open in the editor (the server has not received an open notification
	 * before) the server can send `null` to indicate that the version is
	 * known and the content on disk is the master (as specified with document
	 * content ownership).
	 *
	 * The version number of a document will increase after each change,
	 * including undo/redo. The number doesn't need to be consecutive.
	 */
	Version *Integer `json:"version"`
}

// https://microsoft.github.io/language-server-protocol/specification#textDocumentPositionParams

type TextDocumentPositionParams struct {
	/**
	 * The text document.
	 */
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	/**
	 * The position inside the text document.
	 */
	Position Position `json:"position"`
}

// https://microsoft.github.io/language-server-protocol/specification#documentFilter

type DocumentFilter struct {
	/**
	 * A language id, like `typescript`.
	 */
	Language *string `json:"language,omitempty"`

	/**
	 * A Uri [scheme](#Uri.scheme), like `file` or `untitled`.
	 */
	Scheme *string `json:"scheme,omitempty"`

	/**
	 * A glob pattern, like `*.{ts,js}`.
	 *
	 * Glob patterns can have the following syntax:
	 * - `*` to match one or more characters in a path segment
	 * - `?` to match on one character in a path segment
	 * - `**` to match any number of path segments, including none
	 * - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript
	 *   and JavaScript files)
	 * - `[]` to declare a range of characters to match in a path segment
	 *   (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	 * - `[!...]` to negate a range of characters to match in a path segment
	 *   (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but
	 *   not `example.0`)
	 */
	Pattern *string `json:"pattern,omitempty"`
}

type DocumentSelector []DocumentFilter

// https://microsoft.github.io/language-server-protocol/specification#staticRegistrationOptions

/**
 * Static registration options to be returned in the initialize request.
 */
type StaticRegistrationOptions struct {
	/**
	 * The id used to register the request. The id can be used to deregister
	 * the request again. See also Registration#id.
	 */
	ID *string `json:"id,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specification#textDocumentRegistrationOptions

/**
 * General text document registration options.
 */
type TextDocumentRegistrationOptions struct {
	/**
	 * A document selector to identify the scope of the registration. If set to
	 * null the document selector provided on the client side will be used.
	 */
	DocumentSelector *DocumentSelector `json:"documentSelector"`
}

// https://microsoft.github.io/language-server-protocol/specification#markupContent

/**
 * Describes the content type that a client supports in various
 * result literals like `Hover`, `ParameterInfo` or `CompletionItem`.
 *
 * Please note that `MarkupKinds` must not start with a `$`. This kinds
 * are reserved for internal usage.
 */
type MarkupKind string

const (
	/**
	 * Plain text is supported as a content format
	 */
	MarkupKindPlainText = MarkupKind("plaintext")

	/**
	 * Markdown is supported as a content format
	 */
	MarkupKindMarkdown = MarkupKind("markdown")
)

/**
 * A `MarkupContent` literal represents a string value which content is
 * interpreted base on its kind flag. Currently the protocol supports
 * `plaintext` and `markdown` as markup kinds.
 *
 * If the kind is `markdown` then the value can contain fenced code blocks like
 * in GitHub issues.
 *
 * Here is an example how such a string can be constructed using
 * JavaScript / TypeScript:
 * ```typescript
 * let markdown: MarkdownContent = {
 *  kind: MarkupKind.Markdown,
 *	value: [
 *		'# Header',
 *		'Some text',
 *		'```typescript',
 *		'someCode();',
 *		'```'
 *	].join('\n')
 * };
 * ```
 *
 * *Please Note* that clients might sanitize the return markdown. A client could
 * decide to remove HTML from the markdown to avoid script execution.
 */
type MarkupContent struct {
	/**
	 * The type of the Markup
	 */
	Kind MarkupKind `json:"kind"`

	/**
	 * The content itself
	 */
	Value string `json:"value"`
}

/**
 * Client capabilities specific to the used markdown parser.
 *
 * @since 3.16.0
 */
type MarkdownClientCapabilities struct {
	/**
	 * The name of the parser.
	 */
	Parser string `json:"parser"`

	/**
	 * The version of the parser.
	 */
	Version *string `json:"version,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specification#workDoneProgress

type WorkDoneProgressBegin struct {
	Kind string `json:"kind"` // == "begin"

	/**
	 * Mandatory title of the progress operation. Used to briefly inform about
	 * the kind of operation being performed.
	 *
	 * Examples: "Indexing" or "Linking dependencies".
	 */
	Title string `json:"title"`

	/**
	 * Controls if a cancel button should show to allow the user to cancel the
	 * long running operation. Clients that don't support cancellation are
	 * allowed to ignore the setting.
	 */
	Cancellable *bool `json:"cancellable"`

	/**
	 * Optional, more detailed associated progress message. Contains
	 * complementary information to the `title`.
	 *
	 * Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	 * If unset, the previous progress message (if any) is still valid.
	 */
	Message *string `json:"message"`

	/**
	 * Optional progress percentage to display (value 100 is considered 100%).
	 * If not provided infinite progress is assumed and clients are allowed
	 * to ignore the `percentage` value in subsequent in report notifications.
	 *
	 * The value should be steadily rising. Clients are free to ignore values
	 * that are not following this rule. The value range is [0, 100]
	 */
	Percentage *UInteger `json:"percentage"`
}

type WorkDoneProgressReport struct {
	Kind string `json:"kind"` // == "report"

	/**
	 * Controls enablement state of a cancel button. This property is only valid
	 *  if a cancel button got requested in the `WorkDoneProgressStart` payload.
	 *
	 * Clients that don't support cancellation or don't support control the
	 * button's enablement state are allowed to ignore the setting.
	 */
	Cancellable *bool `json:"cancellable,omitempty"`

	/**
	 * Optional, more detailed associated progress message. Contains
	 * complementary information to the `title`.
	 *
	 * Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	 * If unset, the previous progress message (if any) is still valid.
	 */
	Message *string `json:"message,omitempty"`

	/**
	 * Optional progress percentage to display (value 100 is considered 100%).
	 * If not provided infinite progress is assumed and clients are allowed
	 * to ignore the `percentage` value in subsequent in report notifications.
	 *
	 * The value should be steadily rising. Clients are free to ignore values
	 * that are not following this rule. The value range is [0, 100]
	 */
	Percentage *UInteger `json:"percentage,omitempty"`
}

type WorkDoneProgressEnd struct {
	Kind string `json:"kind"` // == "end"

	/**
	 * Optional, a final message indicating to for example indicate the outcome
	 * of the operation.
	 */
	Message *string `json:"message,omitempty"`
}

type WorkDoneProgressParams struct {
	/**
	 * An optional token that a server can use to report work done progress.
	 */
	WorkDoneToken *ProgressToken `json:"workDoneToken"`
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress *bool `json:"workDoneProgress"`
}

// https://microsoft.github.io/language-server-protocol/specification#partialResults

type PartialResultParams struct {
	/**
	 * An optional token that a server can use to report partial results (e.g.
	 * streaming) to the client.
	 */
	PartialResultToken *ProgressToken `json:"partialResultToken"`
}

// https://microsoft.github.io/language-server-protocol/specification#traceValue

type TraceValue string

const (
	TraceValueOff     = TraceValue("off")
	TraceValueMessage = TraceValue("message")
	TraceValueVerbose = TraceValue("verbose")
)