// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package sesiface provides an interface to enable mocking the Amazon Simple Email Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sesiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ses"
)

// SESAPI provides an interface to enable mocking the
// ses.SES service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Simple Email Service.
//    func myFunc(svc sesiface.SESAPI) bool {
//        // Make svc.CloneReceiptRuleSet request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := ses.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSESClient struct {
//        sesiface.SESAPI
//    }
//    func (m *mockSESClient) CloneReceiptRuleSet(input *ses.CloneReceiptRuleSetInput) (*ses.CloneReceiptRuleSetOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSESClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SESAPI interface {
	CloneReceiptRuleSetRequest(*ses.CloneReceiptRuleSetInput) (*request.Request, *ses.CloneReceiptRuleSetOutput)

	CloneReceiptRuleSet(*ses.CloneReceiptRuleSetInput) (*ses.CloneReceiptRuleSetOutput, error)

	CreateReceiptFilterRequest(*ses.CreateReceiptFilterInput) (*request.Request, *ses.CreateReceiptFilterOutput)

	CreateReceiptFilter(*ses.CreateReceiptFilterInput) (*ses.CreateReceiptFilterOutput, error)

	CreateReceiptRuleRequest(*ses.CreateReceiptRuleInput) (*request.Request, *ses.CreateReceiptRuleOutput)

	CreateReceiptRule(*ses.CreateReceiptRuleInput) (*ses.CreateReceiptRuleOutput, error)

	CreateReceiptRuleSetRequest(*ses.CreateReceiptRuleSetInput) (*request.Request, *ses.CreateReceiptRuleSetOutput)

	CreateReceiptRuleSet(*ses.CreateReceiptRuleSetInput) (*ses.CreateReceiptRuleSetOutput, error)

	DeleteIdentityRequest(*ses.DeleteIdentityInput) (*request.Request, *ses.DeleteIdentityOutput)

	DeleteIdentity(*ses.DeleteIdentityInput) (*ses.DeleteIdentityOutput, error)

	DeleteIdentityPolicyRequest(*ses.DeleteIdentityPolicyInput) (*request.Request, *ses.DeleteIdentityPolicyOutput)

	DeleteIdentityPolicy(*ses.DeleteIdentityPolicyInput) (*ses.DeleteIdentityPolicyOutput, error)

	DeleteReceiptFilterRequest(*ses.DeleteReceiptFilterInput) (*request.Request, *ses.DeleteReceiptFilterOutput)

	DeleteReceiptFilter(*ses.DeleteReceiptFilterInput) (*ses.DeleteReceiptFilterOutput, error)

	DeleteReceiptRuleRequest(*ses.DeleteReceiptRuleInput) (*request.Request, *ses.DeleteReceiptRuleOutput)

	DeleteReceiptRule(*ses.DeleteReceiptRuleInput) (*ses.DeleteReceiptRuleOutput, error)

	DeleteReceiptRuleSetRequest(*ses.DeleteReceiptRuleSetInput) (*request.Request, *ses.DeleteReceiptRuleSetOutput)

	DeleteReceiptRuleSet(*ses.DeleteReceiptRuleSetInput) (*ses.DeleteReceiptRuleSetOutput, error)

	DeleteVerifiedEmailAddressRequest(*ses.DeleteVerifiedEmailAddressInput) (*request.Request, *ses.DeleteVerifiedEmailAddressOutput)

	DeleteVerifiedEmailAddress(*ses.DeleteVerifiedEmailAddressInput) (*ses.DeleteVerifiedEmailAddressOutput, error)

	DescribeActiveReceiptRuleSetRequest(*ses.DescribeActiveReceiptRuleSetInput) (*request.Request, *ses.DescribeActiveReceiptRuleSetOutput)

	DescribeActiveReceiptRuleSet(*ses.DescribeActiveReceiptRuleSetInput) (*ses.DescribeActiveReceiptRuleSetOutput, error)

	DescribeReceiptRuleRequest(*ses.DescribeReceiptRuleInput) (*request.Request, *ses.DescribeReceiptRuleOutput)

	DescribeReceiptRule(*ses.DescribeReceiptRuleInput) (*ses.DescribeReceiptRuleOutput, error)

	DescribeReceiptRuleSetRequest(*ses.DescribeReceiptRuleSetInput) (*request.Request, *ses.DescribeReceiptRuleSetOutput)

	DescribeReceiptRuleSet(*ses.DescribeReceiptRuleSetInput) (*ses.DescribeReceiptRuleSetOutput, error)

	GetIdentityDkimAttributesRequest(*ses.GetIdentityDkimAttributesInput) (*request.Request, *ses.GetIdentityDkimAttributesOutput)

	GetIdentityDkimAttributes(*ses.GetIdentityDkimAttributesInput) (*ses.GetIdentityDkimAttributesOutput, error)

	GetIdentityMailFromDomainAttributesRequest(*ses.GetIdentityMailFromDomainAttributesInput) (*request.Request, *ses.GetIdentityMailFromDomainAttributesOutput)

	GetIdentityMailFromDomainAttributes(*ses.GetIdentityMailFromDomainAttributesInput) (*ses.GetIdentityMailFromDomainAttributesOutput, error)

	GetIdentityNotificationAttributesRequest(*ses.GetIdentityNotificationAttributesInput) (*request.Request, *ses.GetIdentityNotificationAttributesOutput)

	GetIdentityNotificationAttributes(*ses.GetIdentityNotificationAttributesInput) (*ses.GetIdentityNotificationAttributesOutput, error)

	GetIdentityPoliciesRequest(*ses.GetIdentityPoliciesInput) (*request.Request, *ses.GetIdentityPoliciesOutput)

	GetIdentityPolicies(*ses.GetIdentityPoliciesInput) (*ses.GetIdentityPoliciesOutput, error)

	GetIdentityVerificationAttributesRequest(*ses.GetIdentityVerificationAttributesInput) (*request.Request, *ses.GetIdentityVerificationAttributesOutput)

	GetIdentityVerificationAttributes(*ses.GetIdentityVerificationAttributesInput) (*ses.GetIdentityVerificationAttributesOutput, error)

	GetSendQuotaRequest(*ses.GetSendQuotaInput) (*request.Request, *ses.GetSendQuotaOutput)

	GetSendQuota(*ses.GetSendQuotaInput) (*ses.GetSendQuotaOutput, error)

	GetSendStatisticsRequest(*ses.GetSendStatisticsInput) (*request.Request, *ses.GetSendStatisticsOutput)

	GetSendStatistics(*ses.GetSendStatisticsInput) (*ses.GetSendStatisticsOutput, error)

	ListIdentitiesRequest(*ses.ListIdentitiesInput) (*request.Request, *ses.ListIdentitiesOutput)

	ListIdentities(*ses.ListIdentitiesInput) (*ses.ListIdentitiesOutput, error)

	ListIdentitiesPages(*ses.ListIdentitiesInput, func(*ses.ListIdentitiesOutput, bool) bool) error

	ListIdentityPoliciesRequest(*ses.ListIdentityPoliciesInput) (*request.Request, *ses.ListIdentityPoliciesOutput)

	ListIdentityPolicies(*ses.ListIdentityPoliciesInput) (*ses.ListIdentityPoliciesOutput, error)

	ListReceiptFiltersRequest(*ses.ListReceiptFiltersInput) (*request.Request, *ses.ListReceiptFiltersOutput)

	ListReceiptFilters(*ses.ListReceiptFiltersInput) (*ses.ListReceiptFiltersOutput, error)

	ListReceiptRuleSetsRequest(*ses.ListReceiptRuleSetsInput) (*request.Request, *ses.ListReceiptRuleSetsOutput)

	ListReceiptRuleSets(*ses.ListReceiptRuleSetsInput) (*ses.ListReceiptRuleSetsOutput, error)

	ListVerifiedEmailAddressesRequest(*ses.ListVerifiedEmailAddressesInput) (*request.Request, *ses.ListVerifiedEmailAddressesOutput)

	ListVerifiedEmailAddresses(*ses.ListVerifiedEmailAddressesInput) (*ses.ListVerifiedEmailAddressesOutput, error)

	PutIdentityPolicyRequest(*ses.PutIdentityPolicyInput) (*request.Request, *ses.PutIdentityPolicyOutput)

	PutIdentityPolicy(*ses.PutIdentityPolicyInput) (*ses.PutIdentityPolicyOutput, error)

	ReorderReceiptRuleSetRequest(*ses.ReorderReceiptRuleSetInput) (*request.Request, *ses.ReorderReceiptRuleSetOutput)

	ReorderReceiptRuleSet(*ses.ReorderReceiptRuleSetInput) (*ses.ReorderReceiptRuleSetOutput, error)

	SendBounceRequest(*ses.SendBounceInput) (*request.Request, *ses.SendBounceOutput)

	SendBounce(*ses.SendBounceInput) (*ses.SendBounceOutput, error)

	SendEmailRequest(*ses.SendEmailInput) (*request.Request, *ses.SendEmailOutput)

	SendEmail(*ses.SendEmailInput) (*ses.SendEmailOutput, error)

	SendRawEmailRequest(*ses.SendRawEmailInput) (*request.Request, *ses.SendRawEmailOutput)

	SendRawEmail(*ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error)

	SetActiveReceiptRuleSetRequest(*ses.SetActiveReceiptRuleSetInput) (*request.Request, *ses.SetActiveReceiptRuleSetOutput)

	SetActiveReceiptRuleSet(*ses.SetActiveReceiptRuleSetInput) (*ses.SetActiveReceiptRuleSetOutput, error)

	SetIdentityDkimEnabledRequest(*ses.SetIdentityDkimEnabledInput) (*request.Request, *ses.SetIdentityDkimEnabledOutput)

	SetIdentityDkimEnabled(*ses.SetIdentityDkimEnabledInput) (*ses.SetIdentityDkimEnabledOutput, error)

	SetIdentityFeedbackForwardingEnabledRequest(*ses.SetIdentityFeedbackForwardingEnabledInput) (*request.Request, *ses.SetIdentityFeedbackForwardingEnabledOutput)

	SetIdentityFeedbackForwardingEnabled(*ses.SetIdentityFeedbackForwardingEnabledInput) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error)

	SetIdentityHeadersInNotificationsEnabledRequest(*ses.SetIdentityHeadersInNotificationsEnabledInput) (*request.Request, *ses.SetIdentityHeadersInNotificationsEnabledOutput)

	SetIdentityHeadersInNotificationsEnabled(*ses.SetIdentityHeadersInNotificationsEnabledInput) (*ses.SetIdentityHeadersInNotificationsEnabledOutput, error)

	SetIdentityMailFromDomainRequest(*ses.SetIdentityMailFromDomainInput) (*request.Request, *ses.SetIdentityMailFromDomainOutput)

	SetIdentityMailFromDomain(*ses.SetIdentityMailFromDomainInput) (*ses.SetIdentityMailFromDomainOutput, error)

	SetIdentityNotificationTopicRequest(*ses.SetIdentityNotificationTopicInput) (*request.Request, *ses.SetIdentityNotificationTopicOutput)

	SetIdentityNotificationTopic(*ses.SetIdentityNotificationTopicInput) (*ses.SetIdentityNotificationTopicOutput, error)

	SetReceiptRulePositionRequest(*ses.SetReceiptRulePositionInput) (*request.Request, *ses.SetReceiptRulePositionOutput)

	SetReceiptRulePosition(*ses.SetReceiptRulePositionInput) (*ses.SetReceiptRulePositionOutput, error)

	UpdateReceiptRuleRequest(*ses.UpdateReceiptRuleInput) (*request.Request, *ses.UpdateReceiptRuleOutput)

	UpdateReceiptRule(*ses.UpdateReceiptRuleInput) (*ses.UpdateReceiptRuleOutput, error)

	VerifyDomainDkimRequest(*ses.VerifyDomainDkimInput) (*request.Request, *ses.VerifyDomainDkimOutput)

	VerifyDomainDkim(*ses.VerifyDomainDkimInput) (*ses.VerifyDomainDkimOutput, error)

	VerifyDomainIdentityRequest(*ses.VerifyDomainIdentityInput) (*request.Request, *ses.VerifyDomainIdentityOutput)

	VerifyDomainIdentity(*ses.VerifyDomainIdentityInput) (*ses.VerifyDomainIdentityOutput, error)

	VerifyEmailAddressRequest(*ses.VerifyEmailAddressInput) (*request.Request, *ses.VerifyEmailAddressOutput)

	VerifyEmailAddress(*ses.VerifyEmailAddressInput) (*ses.VerifyEmailAddressOutput, error)

	VerifyEmailIdentityRequest(*ses.VerifyEmailIdentityInput) (*request.Request, *ses.VerifyEmailIdentityOutput)

	VerifyEmailIdentity(*ses.VerifyEmailIdentityInput) (*ses.VerifyEmailIdentityOutput, error)

	WaitUntilIdentityExists(*ses.GetIdentityVerificationAttributesInput) error
}

var _ SESAPI = (*ses.SES)(nil)
