package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tags interface{} `json:"Tags"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ResolverRuleConfig represents the ResolverRuleConfig schema from the OpenAPI specification
type ResolverRuleConfig struct {
	Name interface{} `json:"Name,omitempty"`
	Resolverendpointid interface{} `json:"ResolverEndpointId,omitempty"`
	Targetips interface{} `json:"TargetIps,omitempty"`
}

// DisassociateResolverEndpointIpAddressResponse represents the DisassociateResolverEndpointIpAddressResponse schema from the OpenAPI specification
type DisassociateResolverEndpointIpAddressResponse struct {
	Resolverendpoint interface{} `json:"ResolverEndpoint,omitempty"`
}

// CreateFirewallDomainListRequest represents the CreateFirewallDomainListRequest schema from the OpenAPI specification
type CreateFirewallDomainListRequest struct {
	Creatorrequestid interface{} `json:"CreatorRequestId"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DeleteFirewallRuleGroupRequest represents the DeleteFirewallRuleGroupRequest schema from the OpenAPI specification
type DeleteFirewallRuleGroupRequest struct {
	Firewallrulegroupid interface{} `json:"FirewallRuleGroupId"`
}

// ListFirewallRuleGroupsRequest represents the ListFirewallRuleGroupsRequest schema from the OpenAPI specification
type ListFirewallRuleGroupsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AssociateResolverRuleResponse represents the AssociateResolverRuleResponse schema from the OpenAPI specification
type AssociateResolverRuleResponse struct {
	Resolverruleassociation interface{} `json:"ResolverRuleAssociation,omitempty"`
}

// FirewallDomainListMetadata represents the FirewallDomainListMetadata schema from the OpenAPI specification
type FirewallDomainListMetadata struct {
	Name interface{} `json:"Name,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Managedownername interface{} `json:"ManagedOwnerName,omitempty"`
}

// FirewallRule represents the FirewallRule schema from the OpenAPI specification
type FirewallRule struct {
	Blockoverridettl interface{} `json:"BlockOverrideTtl,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Firewalldomainlistid interface{} `json:"FirewallDomainListId,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Action interface{} `json:"Action,omitempty"`
	Blockoverridednstype interface{} `json:"BlockOverrideDnsType,omitempty"`
	Blockoverridedomain interface{} `json:"BlockOverrideDomain,omitempty"`
	Blockresponse interface{} `json:"BlockResponse,omitempty"`
	Modificationtime interface{} `json:"ModificationTime,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Firewallrulegroupid interface{} `json:"FirewallRuleGroupId,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
}

// CreateResolverQueryLogConfigRequest represents the CreateResolverQueryLogConfigRequest schema from the OpenAPI specification
type CreateResolverQueryLogConfigRequest struct {
	Creatorrequestid interface{} `json:"CreatorRequestId"`
	Destinationarn interface{} `json:"DestinationArn"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
}

// UpdateResolverEndpointResponse represents the UpdateResolverEndpointResponse schema from the OpenAPI specification
type UpdateResolverEndpointResponse struct {
	Resolverendpoint interface{} `json:"ResolverEndpoint,omitempty"`
}

// TargetAddress represents the TargetAddress schema from the OpenAPI specification
type TargetAddress struct {
	Port interface{} `json:"Port,omitempty"`
	Ip interface{} `json:"Ip,omitempty"`
	Ipv6 interface{} `json:"Ipv6,omitempty"`
}

// GetResolverQueryLogConfigAssociationRequest represents the GetResolverQueryLogConfigAssociationRequest schema from the OpenAPI specification
type GetResolverQueryLogConfigAssociationRequest struct {
	Resolverquerylogconfigassociationid interface{} `json:"ResolverQueryLogConfigAssociationId"`
}

// ListOutpostResolversResponse represents the ListOutpostResolversResponse schema from the OpenAPI specification
type ListOutpostResolversResponse struct {
	Outpostresolvers interface{} `json:"OutpostResolvers,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListFirewallDomainsRequest represents the ListFirewallDomainsRequest schema from the OpenAPI specification
type ListFirewallDomainsRequest struct {
	Firewalldomainlistid interface{} `json:"FirewallDomainListId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateFirewallDomainsRequest represents the UpdateFirewallDomainsRequest schema from the OpenAPI specification
type UpdateFirewallDomainsRequest struct {
	Domains interface{} `json:"Domains"`
	Firewalldomainlistid interface{} `json:"FirewallDomainListId"`
	Operation interface{} `json:"Operation"`
}

// FirewallConfig represents the FirewallConfig schema from the OpenAPI specification
type FirewallConfig struct {
	Firewallfailopen interface{} `json:"FirewallFailOpen,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
}

// UpdateOutpostResolverResponse represents the UpdateOutpostResolverResponse schema from the OpenAPI specification
type UpdateOutpostResolverResponse struct {
	Outpostresolver interface{} `json:"OutpostResolver,omitempty"`
}

// DisassociateResolverRuleRequest represents the DisassociateResolverRuleRequest schema from the OpenAPI specification
type DisassociateResolverRuleRequest struct {
	Resolverruleid interface{} `json:"ResolverRuleId"`
	Vpcid interface{} `json:"VPCId"`
}

// ResolverDnssecConfig represents the ResolverDnssecConfig schema from the OpenAPI specification
type ResolverDnssecConfig struct {
	Id interface{} `json:"Id,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Validationstatus interface{} `json:"ValidationStatus,omitempty"`
}

// CreateFirewallRuleGroupRequest represents the CreateFirewallRuleGroupRequest schema from the OpenAPI specification
type CreateFirewallRuleGroupRequest struct {
	Creatorrequestid interface{} `json:"CreatorRequestId"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateOutpostResolverResponse represents the CreateOutpostResolverResponse schema from the OpenAPI specification
type CreateOutpostResolverResponse struct {
	Outpostresolver interface{} `json:"OutpostResolver,omitempty"`
}

// DeleteResolverEndpointResponse represents the DeleteResolverEndpointResponse schema from the OpenAPI specification
type DeleteResolverEndpointResponse struct {
	Resolverendpoint interface{} `json:"ResolverEndpoint,omitempty"`
}

// PutResolverQueryLogConfigPolicyRequest represents the PutResolverQueryLogConfigPolicyRequest schema from the OpenAPI specification
type PutResolverQueryLogConfigPolicyRequest struct {
	Arn interface{} `json:"Arn"`
	Resolverquerylogconfigpolicy interface{} `json:"ResolverQueryLogConfigPolicy"`
}

// GetFirewallDomainListResponse represents the GetFirewallDomainListResponse schema from the OpenAPI specification
type GetFirewallDomainListResponse struct {
	Firewalldomainlist interface{} `json:"FirewallDomainList,omitempty"`
}

// GetFirewallRuleGroupAssociationRequest represents the GetFirewallRuleGroupAssociationRequest schema from the OpenAPI specification
type GetFirewallRuleGroupAssociationRequest struct {
	Firewallrulegroupassociationid interface{} `json:"FirewallRuleGroupAssociationId"`
}

// ListFirewallDomainListsRequest represents the ListFirewallDomainListsRequest schema from the OpenAPI specification
type ListFirewallDomainListsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// FirewallDomainList represents the FirewallDomainList schema from the OpenAPI specification
type FirewallDomainList struct {
	Managedownername interface{} `json:"ManagedOwnerName,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Domaincount interface{} `json:"DomainCount,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Modificationtime interface{} `json:"ModificationTime,omitempty"`
}

// CreateOutpostResolverRequest represents the CreateOutpostResolverRequest schema from the OpenAPI specification
type CreateOutpostResolverRequest struct {
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	Name interface{} `json:"Name"`
	Outpostarn interface{} `json:"OutpostArn"`
	Preferredinstancetype interface{} `json:"PreferredInstanceType"`
	Tags interface{} `json:"Tags,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId"`
}

// CreateResolverEndpointRequest represents the CreateResolverEndpointRequest schema from the OpenAPI specification
type CreateResolverEndpointRequest struct {
	Direction interface{} `json:"Direction"`
	Outpostarn interface{} `json:"OutpostArn,omitempty"`
	Ipaddresses interface{} `json:"IpAddresses"`
	Name interface{} `json:"Name,omitempty"`
	Preferredinstancetype interface{} `json:"PreferredInstanceType,omitempty"`
	Resolverendpointtype interface{} `json:"ResolverEndpointType,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds"`
	Creatorrequestid interface{} `json:"CreatorRequestId"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DisassociateResolverQueryLogConfigRequest represents the DisassociateResolverQueryLogConfigRequest schema from the OpenAPI specification
type DisassociateResolverQueryLogConfigRequest struct {
	Resourceid interface{} `json:"ResourceId"`
	Resolverquerylogconfigid interface{} `json:"ResolverQueryLogConfigId"`
}

// GetResolverDnssecConfigResponse represents the GetResolverDnssecConfigResponse schema from the OpenAPI specification
type GetResolverDnssecConfigResponse struct {
	Resolverdnssecconfig interface{} `json:"ResolverDNSSECConfig,omitempty"`
}

// DeleteFirewallRuleResponse represents the DeleteFirewallRuleResponse schema from the OpenAPI specification
type DeleteFirewallRuleResponse struct {
	Firewallrule interface{} `json:"FirewallRule,omitempty"`
}

// ListResolverQueryLogConfigsResponse represents the ListResolverQueryLogConfigsResponse schema from the OpenAPI specification
type ListResolverQueryLogConfigsResponse struct {
	Totalcount interface{} `json:"TotalCount,omitempty"`
	Totalfilteredcount interface{} `json:"TotalFilteredCount,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resolverquerylogconfigs interface{} `json:"ResolverQueryLogConfigs,omitempty"`
}

// DeleteOutpostResolverRequest represents the DeleteOutpostResolverRequest schema from the OpenAPI specification
type DeleteOutpostResolverRequest struct {
	Id interface{} `json:"Id"`
}

// IpAddressRequest represents the IpAddressRequest schema from the OpenAPI specification
type IpAddressRequest struct {
	Ip interface{} `json:"Ip,omitempty"`
	Ipv6 interface{} `json:"Ipv6,omitempty"`
	Subnetid interface{} `json:"SubnetId"`
}

// GetResolverDnssecConfigRequest represents the GetResolverDnssecConfigRequest schema from the OpenAPI specification
type GetResolverDnssecConfigRequest struct {
	Resourceid interface{} `json:"ResourceId"`
}

// ResolverQueryLogConfigAssociation represents the ResolverQueryLogConfigAssociation schema from the OpenAPI specification
type ResolverQueryLogConfigAssociation struct {
	Creationtime interface{} `json:"CreationTime,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Resolverquerylogconfigid interface{} `json:"ResolverQueryLogConfigId,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// GetResolverConfigResponse represents the GetResolverConfigResponse schema from the OpenAPI specification
type GetResolverConfigResponse struct {
	Resolverconfig interface{} `json:"ResolverConfig,omitempty"`
}

// AssociateResolverQueryLogConfigRequest represents the AssociateResolverQueryLogConfigRequest schema from the OpenAPI specification
type AssociateResolverQueryLogConfigRequest struct {
	Resolverquerylogconfigid interface{} `json:"ResolverQueryLogConfigId"`
	Resourceid interface{} `json:"ResourceId"`
}

// ListResolverQueryLogConfigsRequest represents the ListResolverQueryLogConfigsRequest schema from the OpenAPI specification
type ListResolverQueryLogConfigsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
}

// ListOutpostResolversRequest represents the ListOutpostResolversRequest schema from the OpenAPI specification
type ListOutpostResolversRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Outpostarn interface{} `json:"OutpostArn,omitempty"`
}

// DeleteFirewallRuleRequest represents the DeleteFirewallRuleRequest schema from the OpenAPI specification
type DeleteFirewallRuleRequest struct {
	Firewallrulegroupid interface{} `json:"FirewallRuleGroupId"`
	Firewalldomainlistid interface{} `json:"FirewallDomainListId"`
}

// ListFirewallRulesResponse represents the ListFirewallRulesResponse schema from the OpenAPI specification
type ListFirewallRulesResponse struct {
	Firewallrules interface{} `json:"FirewallRules,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetFirewallRuleGroupRequest represents the GetFirewallRuleGroupRequest schema from the OpenAPI specification
type GetFirewallRuleGroupRequest struct {
	Firewallrulegroupid interface{} `json:"FirewallRuleGroupId"`
}

// GetResolverRuleAssociationResponse represents the GetResolverRuleAssociationResponse schema from the OpenAPI specification
type GetResolverRuleAssociationResponse struct {
	Resolverruleassociation interface{} `json:"ResolverRuleAssociation,omitempty"`
}

// DeleteResolverRuleResponse represents the DeleteResolverRuleResponse schema from the OpenAPI specification
type DeleteResolverRuleResponse struct {
	Resolverrule interface{} `json:"ResolverRule,omitempty"`
}

// ListFirewallRuleGroupsResponse represents the ListFirewallRuleGroupsResponse schema from the OpenAPI specification
type ListFirewallRuleGroupsResponse struct {
	Firewallrulegroups interface{} `json:"FirewallRuleGroups,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetFirewallConfigRequest represents the GetFirewallConfigRequest schema from the OpenAPI specification
type GetFirewallConfigRequest struct {
	Resourceid interface{} `json:"ResourceId"`
}

// UpdateFirewallRuleGroupAssociationRequest represents the UpdateFirewallRuleGroupAssociationRequest schema from the OpenAPI specification
type UpdateFirewallRuleGroupAssociationRequest struct {
	Mutationprotection interface{} `json:"MutationProtection,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Firewallrulegroupassociationid interface{} `json:"FirewallRuleGroupAssociationId"`
}

// ListResolverDnssecConfigsRequest represents the ListResolverDnssecConfigsRequest schema from the OpenAPI specification
type ListResolverDnssecConfigsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// ListFirewallRulesRequest represents the ListFirewallRulesRequest schema from the OpenAPI specification
type ListFirewallRulesRequest struct {
	Firewallrulegroupid interface{} `json:"FirewallRuleGroupId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Action interface{} `json:"Action,omitempty"`
}

// ListResolverQueryLogConfigAssociationsRequest represents the ListResolverQueryLogConfigAssociationsRequest schema from the OpenAPI specification
type ListResolverQueryLogConfigAssociationsRequest struct {
	Sortby interface{} `json:"SortBy,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateFirewallDomainListResponse represents the CreateFirewallDomainListResponse schema from the OpenAPI specification
type CreateFirewallDomainListResponse struct {
	Firewalldomainlist interface{} `json:"FirewallDomainList,omitempty"`
}

// AssociateResolverEndpointIpAddressResponse represents the AssociateResolverEndpointIpAddressResponse schema from the OpenAPI specification
type AssociateResolverEndpointIpAddressResponse struct {
	Resolverendpoint interface{} `json:"ResolverEndpoint,omitempty"`
}

// GetResolverQueryLogConfigPolicyRequest represents the GetResolverQueryLogConfigPolicyRequest schema from the OpenAPI specification
type GetResolverQueryLogConfigPolicyRequest struct {
	Arn interface{} `json:"Arn"`
}

// DeleteFirewallDomainListRequest represents the DeleteFirewallDomainListRequest schema from the OpenAPI specification
type DeleteFirewallDomainListRequest struct {
	Firewalldomainlistid interface{} `json:"FirewallDomainListId"`
}

// ResolverConfig represents the ResolverConfig schema from the OpenAPI specification
type ResolverConfig struct {
	Autodefinedreverse interface{} `json:"AutodefinedReverse,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
}

// UpdateFirewallConfigRequest represents the UpdateFirewallConfigRequest schema from the OpenAPI specification
type UpdateFirewallConfigRequest struct {
	Firewallfailopen interface{} `json:"FirewallFailOpen"`
	Resourceid interface{} `json:"ResourceId"`
}

// ListFirewallRuleGroupAssociationsResponse represents the ListFirewallRuleGroupAssociationsResponse schema from the OpenAPI specification
type ListFirewallRuleGroupAssociationsResponse struct {
	Firewallrulegroupassociations interface{} `json:"FirewallRuleGroupAssociations,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// IpAddressUpdate represents the IpAddressUpdate schema from the OpenAPI specification
type IpAddressUpdate struct {
	Ip interface{} `json:"Ip,omitempty"`
	Ipid interface{} `json:"IpId,omitempty"`
	Ipv6 interface{} `json:"Ipv6,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
}

// GetResolverQueryLogConfigResponse represents the GetResolverQueryLogConfigResponse schema from the OpenAPI specification
type GetResolverQueryLogConfigResponse struct {
	Resolverquerylogconfig interface{} `json:"ResolverQueryLogConfig,omitempty"`
}

// UpdateFirewallRuleResponse represents the UpdateFirewallRuleResponse schema from the OpenAPI specification
type UpdateFirewallRuleResponse struct {
	Firewallrule interface{} `json:"FirewallRule,omitempty"`
}

// ResolverQueryLogConfig represents the ResolverQueryLogConfig schema from the OpenAPI specification
type ResolverQueryLogConfig struct {
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Associationcount interface{} `json:"AssociationCount,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Sharestatus interface{} `json:"ShareStatus,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
}

// ListResolverEndpointIpAddressesResponse represents the ListResolverEndpointIpAddressesResponse schema from the OpenAPI specification
type ListResolverEndpointIpAddressesResponse struct {
	Ipaddresses interface{} `json:"IpAddresses,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tagkeys interface{} `json:"TagKeys"`
}

// ListFirewallRuleGroupAssociationsRequest represents the ListFirewallRuleGroupAssociationsRequest schema from the OpenAPI specification
type ListFirewallRuleGroupAssociationsRequest struct {
	Status interface{} `json:"Status,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Firewallrulegroupid interface{} `json:"FirewallRuleGroupId,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
}

// GetResolverRulePolicyResponse represents the GetResolverRulePolicyResponse schema from the OpenAPI specification
type GetResolverRulePolicyResponse struct {
	Resolverrulepolicy interface{} `json:"ResolverRulePolicy,omitempty"`
}

// UpdateOutpostResolverRequest represents the UpdateOutpostResolverRequest schema from the OpenAPI specification
type UpdateOutpostResolverRequest struct {
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Preferredinstancetype interface{} `json:"PreferredInstanceType,omitempty"`
	Id interface{} `json:"Id"`
}

// DisassociateResolverQueryLogConfigResponse represents the DisassociateResolverQueryLogConfigResponse schema from the OpenAPI specification
type DisassociateResolverQueryLogConfigResponse struct {
	Resolverquerylogconfigassociation interface{} `json:"ResolverQueryLogConfigAssociation,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourcearn interface{} `json:"ResourceArn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// DeleteFirewallDomainListResponse represents the DeleteFirewallDomainListResponse schema from the OpenAPI specification
type DeleteFirewallDomainListResponse struct {
	Firewalldomainlist interface{} `json:"FirewallDomainList,omitempty"`
}

// DisassociateResolverEndpointIpAddressRequest represents the DisassociateResolverEndpointIpAddressRequest schema from the OpenAPI specification
type DisassociateResolverEndpointIpAddressRequest struct {
	Ipaddress interface{} `json:"IpAddress"`
	Resolverendpointid interface{} `json:"ResolverEndpointId"`
}

// Filter represents the Filter schema from the OpenAPI specification
type Filter struct {
	Name interface{} `json:"Name,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// ListResolverRulesRequest represents the ListResolverRulesRequest schema from the OpenAPI specification
type ListResolverRulesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// GetResolverQueryLogConfigPolicyResponse represents the GetResolverQueryLogConfigPolicyResponse schema from the OpenAPI specification
type GetResolverQueryLogConfigPolicyResponse struct {
	Resolverquerylogconfigpolicy interface{} `json:"ResolverQueryLogConfigPolicy,omitempty"`
}

// PutFirewallRuleGroupPolicyResponse represents the PutFirewallRuleGroupPolicyResponse schema from the OpenAPI specification
type PutFirewallRuleGroupPolicyResponse struct {
	Returnvalue interface{} `json:"ReturnValue,omitempty"`
}

// CreateResolverQueryLogConfigResponse represents the CreateResolverQueryLogConfigResponse schema from the OpenAPI specification
type CreateResolverQueryLogConfigResponse struct {
	Resolverquerylogconfig interface{} `json:"ResolverQueryLogConfig,omitempty"`
}

// DeleteResolverEndpointRequest represents the DeleteResolverEndpointRequest schema from the OpenAPI specification
type DeleteResolverEndpointRequest struct {
	Resolverendpointid interface{} `json:"ResolverEndpointId"`
}

// GetResolverConfigRequest represents the GetResolverConfigRequest schema from the OpenAPI specification
type GetResolverConfigRequest struct {
	Resourceid interface{} `json:"ResourceId"`
}

// GetOutpostResolverResponse represents the GetOutpostResolverResponse schema from the OpenAPI specification
type GetOutpostResolverResponse struct {
	Outpostresolver interface{} `json:"OutpostResolver,omitempty"`
}

// FirewallRuleGroupAssociation represents the FirewallRuleGroupAssociation schema from the OpenAPI specification
type FirewallRuleGroupAssociation struct {
	Id interface{} `json:"Id,omitempty"`
	Modificationtime interface{} `json:"ModificationTime,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Firewallrulegroupid interface{} `json:"FirewallRuleGroupId,omitempty"`
	Managedownername interface{} `json:"ManagedOwnerName,omitempty"`
	Mutationprotection interface{} `json:"MutationProtection,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
}

// ListResolverRuleAssociationsResponse represents the ListResolverRuleAssociationsResponse schema from the OpenAPI specification
type ListResolverRuleAssociationsResponse struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resolverruleassociations interface{} `json:"ResolverRuleAssociations,omitempty"`
}

// PutFirewallRuleGroupPolicyRequest represents the PutFirewallRuleGroupPolicyRequest schema from the OpenAPI specification
type PutFirewallRuleGroupPolicyRequest struct {
	Arn interface{} `json:"Arn"`
	Firewallrulegrouppolicy interface{} `json:"FirewallRuleGroupPolicy"`
}

// GetFirewallRuleGroupResponse represents the GetFirewallRuleGroupResponse schema from the OpenAPI specification
type GetFirewallRuleGroupResponse struct {
	Firewallrulegroup interface{} `json:"FirewallRuleGroup,omitempty"`
}

// GetFirewallConfigResponse represents the GetFirewallConfigResponse schema from the OpenAPI specification
type GetFirewallConfigResponse struct {
	Firewallconfig interface{} `json:"FirewallConfig,omitempty"`
}

// ListResolverRulesResponse represents the ListResolverRulesResponse schema from the OpenAPI specification
type ListResolverRulesResponse struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resolverrules interface{} `json:"ResolverRules,omitempty"`
}

// GetResolverRuleAssociationRequest represents the GetResolverRuleAssociationRequest schema from the OpenAPI specification
type GetResolverRuleAssociationRequest struct {
	Resolverruleassociationid interface{} `json:"ResolverRuleAssociationId"`
}

// UpdateResolverConfigRequest represents the UpdateResolverConfigRequest schema from the OpenAPI specification
type UpdateResolverConfigRequest struct {
	Autodefinedreverseflag interface{} `json:"AutodefinedReverseFlag"`
	Resourceid interface{} `json:"ResourceId"`
}

// ListResolverEndpointsResponse represents the ListResolverEndpointsResponse schema from the OpenAPI specification
type ListResolverEndpointsResponse struct {
	Resolverendpoints interface{} `json:"ResolverEndpoints,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateResolverRuleResponse represents the UpdateResolverRuleResponse schema from the OpenAPI specification
type UpdateResolverRuleResponse struct {
	Resolverrule interface{} `json:"ResolverRule,omitempty"`
}

// GetFirewallRuleGroupAssociationResponse represents the GetFirewallRuleGroupAssociationResponse schema from the OpenAPI specification
type GetFirewallRuleGroupAssociationResponse struct {
	Firewallrulegroupassociation interface{} `json:"FirewallRuleGroupAssociation,omitempty"`
}

// DeleteOutpostResolverResponse represents the DeleteOutpostResolverResponse schema from the OpenAPI specification
type DeleteOutpostResolverResponse struct {
	Outpostresolver interface{} `json:"OutpostResolver,omitempty"`
}

// ResolverRule represents the ResolverRule schema from the OpenAPI specification
type ResolverRule struct {
	Modificationtime interface{} `json:"ModificationTime,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Sharestatus interface{} `json:"ShareStatus,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Resolverendpointid interface{} `json:"ResolverEndpointId,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Ruletype interface{} `json:"RuleType,omitempty"`
	Targetips interface{} `json:"TargetIps,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// UpdateIpAddress represents the UpdateIpAddress schema from the OpenAPI specification
type UpdateIpAddress struct {
	Ipid interface{} `json:"IpId"`
	Ipv6 interface{} `json:"Ipv6"`
}

// OutpostResolver represents the OutpostResolver schema from the OpenAPI specification
type OutpostResolver struct {
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Preferredinstancetype interface{} `json:"PreferredInstanceType,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Modificationtime interface{} `json:"ModificationTime,omitempty"`
	Outpostarn interface{} `json:"OutpostArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// ListFirewallConfigsRequest represents the ListFirewallConfigsRequest schema from the OpenAPI specification
type ListFirewallConfigsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteResolverQueryLogConfigRequest represents the DeleteResolverQueryLogConfigRequest schema from the OpenAPI specification
type DeleteResolverQueryLogConfigRequest struct {
	Resolverquerylogconfigid interface{} `json:"ResolverQueryLogConfigId"`
}

// DisassociateResolverRuleResponse represents the DisassociateResolverRuleResponse schema from the OpenAPI specification
type DisassociateResolverRuleResponse struct {
	Resolverruleassociation interface{} `json:"ResolverRuleAssociation,omitempty"`
}

// ImportFirewallDomainsRequest represents the ImportFirewallDomainsRequest schema from the OpenAPI specification
type ImportFirewallDomainsRequest struct {
	Domainfileurl interface{} `json:"DomainFileUrl"`
	Firewalldomainlistid interface{} `json:"FirewallDomainListId"`
	Operation interface{} `json:"Operation"`
}

// ListResolverConfigsRequest represents the ListResolverConfigsRequest schema from the OpenAPI specification
type ListResolverConfigsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DisassociateFirewallRuleGroupResponse represents the DisassociateFirewallRuleGroupResponse schema from the OpenAPI specification
type DisassociateFirewallRuleGroupResponse struct {
	Firewallrulegroupassociation interface{} `json:"FirewallRuleGroupAssociation,omitempty"`
}

// GetResolverQueryLogConfigRequest represents the GetResolverQueryLogConfigRequest schema from the OpenAPI specification
type GetResolverQueryLogConfigRequest struct {
	Resolverquerylogconfigid interface{} `json:"ResolverQueryLogConfigId"`
}

// CreateFirewallRuleResponse represents the CreateFirewallRuleResponse schema from the OpenAPI specification
type CreateFirewallRuleResponse struct {
	Firewallrule interface{} `json:"FirewallRule,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// UpdateFirewallRuleRequest represents the UpdateFirewallRuleRequest schema from the OpenAPI specification
type UpdateFirewallRuleRequest struct {
	Priority interface{} `json:"Priority,omitempty"`
	Action interface{} `json:"Action,omitempty"`
	Blockoverridedomain interface{} `json:"BlockOverrideDomain,omitempty"`
	Blockresponse interface{} `json:"BlockResponse,omitempty"`
	Firewalldomainlistid interface{} `json:"FirewallDomainListId"`
	Firewallrulegroupid interface{} `json:"FirewallRuleGroupId"`
	Name interface{} `json:"Name,omitempty"`
	Blockoverridettl interface{} `json:"BlockOverrideTtl,omitempty"`
	Blockoverridednstype interface{} `json:"BlockOverrideDnsType,omitempty"`
}

// UpdateFirewallConfigResponse represents the UpdateFirewallConfigResponse schema from the OpenAPI specification
type UpdateFirewallConfigResponse struct {
	Firewallconfig interface{} `json:"FirewallConfig,omitempty"`
}

// UpdateResolverRuleRequest represents the UpdateResolverRuleRequest schema from the OpenAPI specification
type UpdateResolverRuleRequest struct {
	Config interface{} `json:"Config"`
	Resolverruleid interface{} `json:"ResolverRuleId"`
}

// CreateFirewallRuleGroupResponse represents the CreateFirewallRuleGroupResponse schema from the OpenAPI specification
type CreateFirewallRuleGroupResponse struct {
	Firewallrulegroup interface{} `json:"FirewallRuleGroup,omitempty"`
}

// DeleteResolverQueryLogConfigResponse represents the DeleteResolverQueryLogConfigResponse schema from the OpenAPI specification
type DeleteResolverQueryLogConfigResponse struct {
	Resolverquerylogconfig interface{} `json:"ResolverQueryLogConfig,omitempty"`
}

// UpdateFirewallDomainsResponse represents the UpdateFirewallDomainsResponse schema from the OpenAPI specification
type UpdateFirewallDomainsResponse struct {
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// GetResolverRuleResponse represents the GetResolverRuleResponse schema from the OpenAPI specification
type GetResolverRuleResponse struct {
	Resolverrule interface{} `json:"ResolverRule,omitempty"`
}

// CreateResolverEndpointResponse represents the CreateResolverEndpointResponse schema from the OpenAPI specification
type CreateResolverEndpointResponse struct {
	Resolverendpoint interface{} `json:"ResolverEndpoint,omitempty"`
}

// GetFirewallRuleGroupPolicyRequest represents the GetFirewallRuleGroupPolicyRequest schema from the OpenAPI specification
type GetFirewallRuleGroupPolicyRequest struct {
	Arn interface{} `json:"Arn"`
}

// ListResolverDnssecConfigsResponse represents the ListResolverDnssecConfigsResponse schema from the OpenAPI specification
type ListResolverDnssecConfigsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resolverdnssecconfigs interface{} `json:"ResolverDnssecConfigs,omitempty"`
}

// GetResolverRuleRequest represents the GetResolverRuleRequest schema from the OpenAPI specification
type GetResolverRuleRequest struct {
	Resolverruleid interface{} `json:"ResolverRuleId"`
}

// GetFirewallDomainListRequest represents the GetFirewallDomainListRequest schema from the OpenAPI specification
type GetFirewallDomainListRequest struct {
	Firewalldomainlistid interface{} `json:"FirewallDomainListId"`
}

// ResolverEndpoint represents the ResolverEndpoint schema from the OpenAPI specification
type ResolverEndpoint struct {
	Modificationtime interface{} `json:"ModificationTime,omitempty"`
	Preferredinstancetype interface{} `json:"PreferredInstanceType,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Hostvpcid interface{} `json:"HostVPCId,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Outpostarn interface{} `json:"OutpostArn,omitempty"`
	Direction interface{} `json:"Direction,omitempty"`
	Resolverendpointtype interface{} `json:"ResolverEndpointType,omitempty"`
	Ipaddresscount interface{} `json:"IpAddressCount,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// GetFirewallRuleGroupPolicyResponse represents the GetFirewallRuleGroupPolicyResponse schema from the OpenAPI specification
type GetFirewallRuleGroupPolicyResponse struct {
	Firewallrulegrouppolicy interface{} `json:"FirewallRuleGroupPolicy,omitempty"`
}

// ListResolverEndpointIpAddressesRequest represents the ListResolverEndpointIpAddressesRequest schema from the OpenAPI specification
type ListResolverEndpointIpAddressesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resolverendpointid interface{} `json:"ResolverEndpointId"`
}

// GetOutpostResolverRequest represents the GetOutpostResolverRequest schema from the OpenAPI specification
type GetOutpostResolverRequest struct {
	Id interface{} `json:"Id"`
}

// GetResolverEndpointRequest represents the GetResolverEndpointRequest schema from the OpenAPI specification
type GetResolverEndpointRequest struct {
	Resolverendpointid interface{} `json:"ResolverEndpointId"`
}

// DeleteFirewallRuleGroupResponse represents the DeleteFirewallRuleGroupResponse schema from the OpenAPI specification
type DeleteFirewallRuleGroupResponse struct {
	Firewallrulegroup interface{} `json:"FirewallRuleGroup,omitempty"`
}

// UpdateResolverDnssecConfigResponse represents the UpdateResolverDnssecConfigResponse schema from the OpenAPI specification
type UpdateResolverDnssecConfigResponse struct {
	Resolverdnssecconfig interface{} `json:"ResolverDNSSECConfig,omitempty"`
}

// FirewallRuleGroupMetadata represents the FirewallRuleGroupMetadata schema from the OpenAPI specification
type FirewallRuleGroupMetadata struct {
	Name interface{} `json:"Name,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Sharestatus interface{} `json:"ShareStatus,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// CreateResolverRuleResponse represents the CreateResolverRuleResponse schema from the OpenAPI specification
type CreateResolverRuleResponse struct {
	Resolverrule interface{} `json:"ResolverRule,omitempty"`
}

// CreateFirewallRuleRequest represents the CreateFirewallRuleRequest schema from the OpenAPI specification
type CreateFirewallRuleRequest struct {
	Blockoverridednstype interface{} `json:"BlockOverrideDnsType,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId"`
	Name interface{} `json:"Name"`
	Priority interface{} `json:"Priority"`
	Action interface{} `json:"Action"`
	Blockoverridedomain interface{} `json:"BlockOverrideDomain,omitempty"`
	Blockoverridettl interface{} `json:"BlockOverrideTtl,omitempty"`
	Blockresponse interface{} `json:"BlockResponse,omitempty"`
	Firewalldomainlistid interface{} `json:"FirewallDomainListId"`
	Firewallrulegroupid interface{} `json:"FirewallRuleGroupId"`
}

// UpdateResolverDnssecConfigRequest represents the UpdateResolverDnssecConfigRequest schema from the OpenAPI specification
type UpdateResolverDnssecConfigRequest struct {
	Validation interface{} `json:"Validation"`
	Resourceid interface{} `json:"ResourceId"`
}

// ResolverRuleAssociation represents the ResolverRuleAssociation schema from the OpenAPI specification
type ResolverRuleAssociation struct {
	Name interface{} `json:"Name,omitempty"`
	Resolverruleid interface{} `json:"ResolverRuleId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Vpcid interface{} `json:"VPCId,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// GetResolverEndpointResponse represents the GetResolverEndpointResponse schema from the OpenAPI specification
type GetResolverEndpointResponse struct {
	Resolverendpoint interface{} `json:"ResolverEndpoint,omitempty"`
}

// AssociateFirewallRuleGroupResponse represents the AssociateFirewallRuleGroupResponse schema from the OpenAPI specification
type AssociateFirewallRuleGroupResponse struct {
	Firewallrulegroupassociation interface{} `json:"FirewallRuleGroupAssociation,omitempty"`
}

// AssociateResolverEndpointIpAddressRequest represents the AssociateResolverEndpointIpAddressRequest schema from the OpenAPI specification
type AssociateResolverEndpointIpAddressRequest struct {
	Ipaddress interface{} `json:"IpAddress"`
	Resolverendpointid interface{} `json:"ResolverEndpointId"`
}

// FirewallRuleGroup represents the FirewallRuleGroup schema from the OpenAPI specification
type FirewallRuleGroup struct {
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Modificationtime interface{} `json:"ModificationTime,omitempty"`
	Sharestatus interface{} `json:"ShareStatus,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Rulecount interface{} `json:"RuleCount,omitempty"`
}

// ListFirewallConfigsResponse represents the ListFirewallConfigsResponse schema from the OpenAPI specification
type ListFirewallConfigsResponse struct {
	Firewallconfigs interface{} `json:"FirewallConfigs,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListFirewallDomainListsResponse represents the ListFirewallDomainListsResponse schema from the OpenAPI specification
type ListFirewallDomainListsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Firewalldomainlists interface{} `json:"FirewallDomainLists,omitempty"`
}

// ImportFirewallDomainsResponse represents the ImportFirewallDomainsResponse schema from the OpenAPI specification
type ImportFirewallDomainsResponse struct {
	Status interface{} `json:"Status,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// DisassociateFirewallRuleGroupRequest represents the DisassociateFirewallRuleGroupRequest schema from the OpenAPI specification
type DisassociateFirewallRuleGroupRequest struct {
	Firewallrulegroupassociationid interface{} `json:"FirewallRuleGroupAssociationId"`
}

// UpdateFirewallRuleGroupAssociationResponse represents the UpdateFirewallRuleGroupAssociationResponse schema from the OpenAPI specification
type UpdateFirewallRuleGroupAssociationResponse struct {
	Firewallrulegroupassociation interface{} `json:"FirewallRuleGroupAssociation,omitempty"`
}

// ListResolverEndpointsRequest represents the ListResolverEndpointsRequest schema from the OpenAPI specification
type ListResolverEndpointsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateResolverEndpointRequest represents the UpdateResolverEndpointRequest schema from the OpenAPI specification
type UpdateResolverEndpointRequest struct {
	Name interface{} `json:"Name,omitempty"`
	Resolverendpointid interface{} `json:"ResolverEndpointId"`
	Resolverendpointtype interface{} `json:"ResolverEndpointType,omitempty"`
	Updateipaddresses interface{} `json:"UpdateIpAddresses,omitempty"`
}

// DeleteResolverRuleRequest represents the DeleteResolverRuleRequest schema from the OpenAPI specification
type DeleteResolverRuleRequest struct {
	Resolverruleid interface{} `json:"ResolverRuleId"`
}

// GetResolverRulePolicyRequest represents the GetResolverRulePolicyRequest schema from the OpenAPI specification
type GetResolverRulePolicyRequest struct {
	Arn interface{} `json:"Arn"`
}

// ListResolverConfigsResponse represents the ListResolverConfigsResponse schema from the OpenAPI specification
type ListResolverConfigsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resolverconfigs interface{} `json:"ResolverConfigs,omitempty"`
}

// CreateResolverRuleRequest represents the CreateResolverRuleRequest schema from the OpenAPI specification
type CreateResolverRuleRequest struct {
	Creatorrequestid interface{} `json:"CreatorRequestId"`
	Domainname interface{} `json:"DomainName"`
	Name interface{} `json:"Name,omitempty"`
	Resolverendpointid interface{} `json:"ResolverEndpointId,omitempty"`
	Ruletype interface{} `json:"RuleType"`
	Tags interface{} `json:"Tags,omitempty"`
	Targetips interface{} `json:"TargetIps,omitempty"`
}

// PutResolverRulePolicyRequest represents the PutResolverRulePolicyRequest schema from the OpenAPI specification
type PutResolverRulePolicyRequest struct {
	Arn interface{} `json:"Arn"`
	Resolverrulepolicy interface{} `json:"ResolverRulePolicy"`
}

// AssociateResolverQueryLogConfigResponse represents the AssociateResolverQueryLogConfigResponse schema from the OpenAPI specification
type AssociateResolverQueryLogConfigResponse struct {
	Resolverquerylogconfigassociation interface{} `json:"ResolverQueryLogConfigAssociation,omitempty"`
}

// IpAddressResponse represents the IpAddressResponse schema from the OpenAPI specification
type IpAddressResponse struct {
	Ipid interface{} `json:"IpId,omitempty"`
	Ipv6 interface{} `json:"Ipv6,omitempty"`
	Modificationtime interface{} `json:"ModificationTime,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Ip interface{} `json:"Ip,omitempty"`
}

// ListFirewallDomainsResponse represents the ListFirewallDomainsResponse schema from the OpenAPI specification
type ListFirewallDomainsResponse struct {
	Domains interface{} `json:"Domains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// PutResolverRulePolicyResponse represents the PutResolverRulePolicyResponse schema from the OpenAPI specification
type PutResolverRulePolicyResponse struct {
	Returnvalue interface{} `json:"ReturnValue,omitempty"`
}

// ListResolverQueryLogConfigAssociationsResponse represents the ListResolverQueryLogConfigAssociationsResponse schema from the OpenAPI specification
type ListResolverQueryLogConfigAssociationsResponse struct {
	Totalfilteredcount interface{} `json:"TotalFilteredCount,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resolverquerylogconfigassociations interface{} `json:"ResolverQueryLogConfigAssociations,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
}

// AssociateFirewallRuleGroupRequest represents the AssociateFirewallRuleGroupRequest schema from the OpenAPI specification
type AssociateFirewallRuleGroupRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Vpcid interface{} `json:"VpcId"`
	Creatorrequestid interface{} `json:"CreatorRequestId"`
	Firewallrulegroupid interface{} `json:"FirewallRuleGroupId"`
	Mutationprotection interface{} `json:"MutationProtection,omitempty"`
	Name interface{} `json:"Name"`
	Priority interface{} `json:"Priority"`
}

// PutResolverQueryLogConfigPolicyResponse represents the PutResolverQueryLogConfigPolicyResponse schema from the OpenAPI specification
type PutResolverQueryLogConfigPolicyResponse struct {
	Returnvalue interface{} `json:"ReturnValue,omitempty"`
}

// GetResolverQueryLogConfigAssociationResponse represents the GetResolverQueryLogConfigAssociationResponse schema from the OpenAPI specification
type GetResolverQueryLogConfigAssociationResponse struct {
	Resolverquerylogconfigassociation interface{} `json:"ResolverQueryLogConfigAssociation,omitempty"`
}

// ListResolverRuleAssociationsRequest represents the ListResolverRuleAssociationsRequest schema from the OpenAPI specification
type ListResolverRuleAssociationsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// UpdateResolverConfigResponse represents the UpdateResolverConfigResponse schema from the OpenAPI specification
type UpdateResolverConfigResponse struct {
	Resolverconfig interface{} `json:"ResolverConfig,omitempty"`
}

// AssociateResolverRuleRequest represents the AssociateResolverRuleRequest schema from the OpenAPI specification
type AssociateResolverRuleRequest struct {
	Name interface{} `json:"Name,omitempty"`
	Resolverruleid interface{} `json:"ResolverRuleId"`
	Vpcid interface{} `json:"VPCId"`
}
