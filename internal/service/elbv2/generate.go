//go:generate go run -tags generate ../../generate/tags/main.go -ListTags=yes -ListTagsOp=DescribeTags -ListTagsInIDElem=ResourceArns -ListTagsInIDNeedSlice=yes -ListTagsOutTagsElem=TagDescriptions[0].Tags -ServiceTagsSlice=yes -TagOp=AddTags -TagInIDElem=ResourceArns -TagInIDNeedSlice=yes -UntagOp=RemoveTags -UpdateTags=yes
// ONLY generate directives and package declaration! Do not add anything else to this file.

package elbv2
