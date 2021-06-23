/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing, and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// ImageViewDetail struct for ImageViewDetail
type ImageViewDetail struct {
	// ImageViewDetail ID
	ID string `json:"ID,omitempty"`
	// ImageIndicator is a code that indicates the presence and disposition of an image view conveyed in the related ImageViewData.  When an image view is not present (0) then certain conditional fields in this ImageViewDetail and the related ImageViewData shall not be present and will be filled with blank space.  * `0` - Image view not present * `1` - Image view present, actual check * `2` - Image view present, not actual check * `3` - Image view present, unable to determine if value is 1 or 2
	ImageIndicator int32 `json:"imageIndicator,omitempty"`
	// ImageCreatorRoutingNumber identifies the financial institution that created the image view.
	ImageCreatorRoutingNumber string `json:"imageCreatorRoutingNumber,omitempty"`
	// ImageCreatorDate is the date assigned by the image creator for the image view conveyed in the related ImageData.
	ImageCreatorDate time.Time `json:"imageCreatorDate,omitempty"`
	// ImageViewFormatIndicator is a code that identifies the type of image format used in the related ImageViewData.ImageData. The image format type is also commonly specified by reference to the file extension used when image data is saved as an image file.  Agreement not required: * `00` - TIFF 6; Extension: TIF  Agreement required: * `01` - IOCA FS 11; Extension: ICA * `20` - PNG (Portable Network Graphics); Extension: PNG ‘21’ JFIF (JPEG File Interchange Format); Extension: JPG * `22` - SPIFF (Still Picture Interchange File Format) (ITU-T Rec. T.84 Annex F); Extension: SPF * `23` - JBIG data stream (ITU-T Rec. T.82/ISO/IEC 11544:1993); Extension: JBG ‘24’ JPEG 2000 (ISO/IEC 15444-1:2000); Extension: JP2
	ImageViewFormatIndicator string `json:"imageViewFormatIndicator,omitempty"`
	// ImageViewCompressionAlgorithm is a code that identifies the algorithm or method used to compress the Image Data in the related ImageViewData.ImageData.  Agreement not required: * `00` - Group 4 facsimile compression (ITU-T Rec. T.563/CCITT Rec. T.6)  Agreement required: * `01` - JPEG Baseline (JPEG Interchange Format) (ITU-T Rec. T.81/ISO/IEC 10918) * `02` - ABIC * `21` - PNG (Portable Network Graphics) * `22` - JBIG (ITU-T Rec. T.82/ISO/IEC 11544:1993) * `23` - JPEG 2000 (ISO/IEC 15444–1:2000)
	ImageViewCompressionAlgorithm string `json:"imageViewCompressionAlgorithm,omitempty"`
	// ImageViewDataSize is the total number of bytes in ImageViewData.ImageData.  Use of this field is NOT recommended. If data is present it shall be ignored, and ImageViewData.ImageDataLength shall take precedence.
	ImageViewDataSize string `json:"imageViewDataSize,omitempty"`
	// ViewSideIndicator is a code that indicates the image view conveyed in the related ImageViewData. An image view may be a full view of the item (i.e., the entire full face of the document) or may be a partial view (snippet) as determined by viewDescriptor.  * `0` - Front image view * `1` - Rear image view
	ViewSideIndicator int32 `json:"viewSideIndicator,omitempty"`
	// ViewDescriptor is a code that indicates the nature of the image view based on ImageViewData.ImageData.  * `00` -  Full view * `01` -  Partial view–unspecified Area of Interest * `02` -  Partial view–date Area of Interest * `03` -  Partial view–payee Area of Interest * `04` -  Partial view–convenience amount Area of Interest * `05` -  Partial view–amount in words (legal amount) Area of Interest * `06` -  Partial view–signature Area(s) of Interest * `07` -  Partial view–payor name and address Area of Interest * `08` -  Partial view–MICR line Area of Interest * `09` -  Partial view–memo line Area of Interest * `10` -  Partial view–payor bank name and address Area of Interest * `11` -  Partial view–payee endorsement Area of Interest * `12` -  Partial view–Bank Of First Deposit (BOFD) endorsement Area of Interest * `13` -  Partial view–transit endorsement Area of Interest * `14 - 99` -  Reserved for X9
	ViewDescriptor string `json:"viewDescriptor,omitempty"`
	// DigitalSignatureIndicator is a code that indicates the presence or absence of a digital signature for the image view contained in ImageViewData.ImageData. If present, the Digital Signature is conveyed in the related DigitalSignature.  * `0` - Digital Signature is not present * `1` - Digital Signature is present
	DigitalSignatureIndicator int32 `json:"digitalSignatureIndicator,omitempty"`
	// DigitalSignatureMethod is a code that identifies the cryptographic algorithm used to generate and validate the Digital Signature in ImageViewData.DigitalSignature. * `00` - Digital Signature Algorithm (DSA) with SHA1 (ANSI X9.30) * `01` - RSA with MD5 (ANSI X9.31) * `02` - RSA with MDC2 (ANSI X9.31) * `03` - RSA with SHA1 (ANSI X9.31) * `04` - Elliptic Curve DSA (ECDSA) with SHA1 (ANSI X9.62) * `05 - 99` - Reserved for emerging cryptographic algorithms.
	DigitalSignatureMethod string `json:"digitalSignatureMethod,omitempty"`
	// SecurityKeySize is the length in bits of the cryptographic algorithm key used to create the Digital Signature. Valid values 0–99999
	SecurityKeySize int32 `json:"securityKeySize,omitempty"`
	// ProtectedDataStart is a number that represents the offset in bytes from the first byte (counted as byte 1) of the image data in ImageViewData.ImageData to the first byte of the image data protected by the digital signature.  * 0000000 - Digital Signature is applied to entire image data * 000001–9999999 - Valid offset values
	ProtectedDataStart int32 `json:"protectedDataStart,omitempty"`
	// ProtectedDataLength is a number of contiguous bytes of image data in the related ImageViewData.ImageData protected by the digital signature starting with the byte indicated by the value of the ProtectedDataStart in this ImageViewDetail. The ProtectedDataLength value shall not exceed the ImageViewData.ImageDataLength.  * 0000000 - Digital Signature is applied to entire image data * 000001–9999999 - Valid length values
	ProtectedDataLength int32 `json:"protectedDataLength,omitempty"`
	// ImageRecreateIndicator is a code that indicates whether the sender has the ability to recreate the image view conveyed in the related ImageViewData.ImageData.  * `0` - Sender can recreate the image view for the duration of the agreed upon retention time frames. * `1` - Sender cannot recreate image view.
	ImageRecreateIndicator int32 `json:"imageRecreateIndicator,omitempty"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField,omitempty"`
	// OverrideIndicator is a code that indicates to a receiving exchange partner that this image view has a detected image test failure that cannot be corrected and that this view shall be accepted regardless of any image test failures.  * ` ` -  blank/space indicates no observed image test failure present * `0` -  No override information for this view or not applicable * `1` -  Imperfect image * `A` -  IQA Fail–Image view reviewed and deemed usable—no alternate format * `B` -  IQA Fail–Image view reviewed and deemed usable—alternate format included in this file * `C` -  IQA Fail–Image view reviewed and deemed usable–alternate format included in this file and original document available * `D` -  IQA Fail–Image view reviewed and deemed usable–alternate format available * `E` -  IQA Fail–Image view reviewed and deemed usable–original document available * `F` -  IQA Fail–Image view reviewed and deemed usable–original document and alternate format available * `G` -  IQA Fail–Image view reviewed and deemed unusable–no alternate format * `H` -  IQA Fail–Image view reviewed and deemed unusable–alternate format included in this file * `I` -  IQA Fail–Image view reviewed and deemed unusable–alternate format included in this file and original document available * `J` -  IQA Fail–Image view reviewed and deemed unusable–alternate format available * `K` -  IQA Fail–Image view reviewed and deemed unusable–original document available * `L` -  IQA Fail–Image view reviewed and deemed unusable–original document and alternate format available * `M` -  IQA Fail–Image view not reviewed–no alternate format * `N` -  IQA Fail–Image view not reviewed–alternate format included in this file * `O` -  IQA Fail–Image view not reviewed–alternate format included in this file and original
	OverrideIndicator string `json:"overrideIndicator,omitempty"`
}
