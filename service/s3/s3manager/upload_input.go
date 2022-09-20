// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3manager

import (
	"io"
	"time"
)

// UploadInput provides the input parameters for uploading a stream or buffer
// to an object in an Amazon S3 bucket. This type is similar to the s3
// package's PutObjectInput with the exception that the Body member is an
// io.Reader instead of an io.ReadSeeker.
//
// The ContentMD5 member for pre-computed MD5 checksums will be ignored for
// multipart uploads. Objects that will be uploaded in a single part, the
// ContentMD5 will be used.
//
// The Checksum members for pre-computed checksums will be ignored for
// multipart uploads. Objects that will be uploaded in a single part, will
// include the checksum member in the request.
type UploadInput struct {
	_ struct{} `locationName:"PutObjectRequest" type:"structure" payload:"Body"`

	// The canned ACL to apply to the object. For more information, see Canned ACL
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
	//
	// This action is not supported by Amazon S3 on Outposts.
	ACL *string `location:"header" locationName:"x-amz-acl" type:"string" enum:"ObjectCannedACL"`

	// The readable body payload to send to S3.
	Body io.Reader

	// The bucket name to which the PUT action was initiated.
	//
	// When using this action with an access point, you must direct requests to
	// the access point hostname. The access point hostname takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this action with an access point through the Amazon Web Services
	// SDKs, you provide the access point ARN in place of the bucket name. For more
	// information about access point ARNs, see Using access points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide.
	//
	// When using this action with Amazon S3 on Outposts, you must direct requests
	// to the S3 on Outposts hostname. The S3 on Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com. When
	// using this action with S3 on Outposts through the Amazon Web Services SDKs,
	// you provide the Outposts bucket ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see Using Amazon S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
	// in the Amazon S3 User Guide.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Can be used to specify caching behavior along the request/reply chain. For
	// more information, see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9).
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// Specifies presentational information for the object. For more information,
	// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1 (http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1).
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field. For more information, see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11).
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`

	// The base64-encoded 128-bit MD5 digest of the message (without the headers)
	// according to RFC 1864. This header can be used as a message integrity check
	// to verify that the data is the same data that was originally sent. Although
	// it is optional, we recommend using the Content-MD5 mechanism as an end-to-end
	// integrity check. For more information about REST request authentication,
	// see REST Authentication (https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html).
	//
	// If the ContentMD5 is provided for a multipart upload, it will be ignored.
	// Objects that will be uploaded in a single part, the ContentMD5 will be used.
	ContentMD5 *string `location:"header" locationName:"Content-MD5" type:"string"`

	// A standard MIME type describing the format of the contents. For more information,
	// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17 (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	//ExpectedBucketOwner *string `location:"header" locationName:"x-amz-expected-bucket-owner" type:"string"`

	// The date and time at which the object is no longer cacheable. For more information,
	// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21 (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21).
	Expires *time.Time `location:"header" locationName:"Expires" type:"timestamp"`

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	//
	// This action is not supported by Amazon S3 on Outposts.
	GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`

	// Allows grantee to read the object data and its metadata.
	//
	// This action is not supported by Amazon S3 on Outposts.
	GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`

	// Allows grantee to read the object ACL.
	//
	// This action is not supported by Amazon S3 on Outposts.
	GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`

	// Allows grantee to write the ACL for the applicable object.
	//
	// This action is not supported by Amazon S3 on Outposts.
	GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`

	// Object key for which the PUT action was initiated.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// A map of metadata to store with the object in S3.
	Metadata map[string]*string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer *string `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"RequestPayer"`

	// Date on which it will be legal to delete or modify the object. This field
	// can only be specified if Retention-Directive is REPLACE. You can only specify
	// this or the Retention-Period header. If both are specified a 400 error will
	// be returned. If neither is specified the bucket's DefaultRetention period
	// will be used.
	RetentionExpirationDate *time.Time `location:"header" locationName:"Retention-Expiration-Date" type:"timestamp"`

	// A single legal hold to apply to the object. This field can only be specified
	// if Retention-Directive is REPLACE. A legal hold is a character long string
	// of max length 64. The object cannot be overwritten or deleted until all legal
	// holds associated with the object are removed.
	RetentionLegalHoldId *string `location:"header" locationName:"Retention-Legal-Hold-ID" type:"string"`

	// Retention period to store on the object in seconds. If this field and Retention-Expiration-Date
	// are specified a 400 error is returned. If neither is specified the bucket's
	// DefaultRetention period will be used. 0 is a legal value assuming the bucket's
	// minimum retention period is also 0.
	RetentionPeriod *int64 `location:"header" locationName:"Retention-Period" type:"integer"`

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// S3 does not store the encryption key. The key must be appropriate for use
	// with the algorithm specified in the x-amz-server-side-encryption-customer-algorithm
	// header.
	SSECustomerKey *string `marshal-as:"blob" location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string" sensitive:"true"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If x-amz-server-side-encryption is present and has the value of aws:kms,
	// this header specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetrical customer managed customer master key (CMK) that was used for
	// the object.
	//
	// If the value of x-amz-server-side-encryption is aws:kms, this header specifies
	// the ID of the symmetric customer managed AWS KMS CMK that will be used for
	// the object. If you specify x-amz-server-side-encryption:aws:kms, but do not
	// providex-amz-server-side-encryption-aws-kms-key-id, Amazon S3 uses the AWS
	// managed CMK in AWS to protect the data.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The server-side encryption algorithm used when storing this object in Amazon
	// S3 (for example, AES256, aws:kms).
	ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`

	// By default, Amazon S3 uses the STANDARD Storage Class to store newly created
	// objects. The STANDARD storage class provides high durability and high availability.
	// Depending on performance needs, you can specify a different Storage Class.
	// Amazon S3 on Outposts only uses the OUTPOSTS Storage Class. For more information,
	// see Storage Classes (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
	// in the Amazon S3 Service Developer Guide.
	StorageClass *string `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"StorageClass"`

	// The tag-set for the object. The tag-set must be encoded as URL Query parameters.
	// (For example, "Key1=Value1")
	Tagging *string `location:"header" locationName:"x-amz-tagging" type:"string"`

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata. For information about object
	// metadata, see Object Key and Metadata (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html).
	//
	// In the following example, the request header sets the redirect to an object
	// (anotherPage.html) in the same bucket:
	//
	// x-amz-website-redirect-location: /anotherPage.html
	//
	// In the following example, the request header sets the object redirect to
	// another website:
	//
	// x-amz-website-redirect-location: http://www.example.com/
	//
	// For more information about website hosting in Amazon S3, see Hosting Websites
	// on Amazon S3 (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html)
	// and How to Configure Website Page Redirects (https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}
