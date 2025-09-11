# TMForum Object Validation Report

**Generated:** 2025-09-11 05:42:32 UTC

**Configuration:**
- Base URL: `https://tmf.dome-marketplace-sbx.org`
- Object Types: billPresentationMedia, cancelResourceOrder, customer, partyRole, resource, resourceSpecification, scale, serviceCatalog, individual, partyAccount, resourceCatalog, resourceOrder, serviceSpecification, category, customerBill, financialAccount, heal, organization, product, productOffering, productOfferingPrice, agreement, billFormat, billingCycleSpecification, cancelProductOrder, productSpecification, quote, resourceFunction, serviceCandidate, customerBillOnDemand, settlementAccount, agreementSpecification, monitor, resourceCandidate, serviceCategory, usageSpecification, billingAccount, service, appliedCustomerBillingRate, cancelServiceOrder, catalog, migrate, productOrder, resourceCategory, serviceOrder, usage
- Timeout: 30 seconds
- Validate Required Fields: true
- Validate Related Party: true

## Summary Statistics

| Metric | Value |
|--------|-------|
| Total Objects | 515 |
| Valid Objects | 0 |
| Invalid Objects | 515 |
| Total Errors | 929 |
| Total Warnings | 0 |
| Processing Time | 77.245µs |

## Statistics by Object Type

| Object Type | Count | Valid | Invalid | Errors | Warnings |
|-------------|-------|-------|---------|--------|----------|
| appliedCustomerBillingRate | 6 | 0 | 6 | 27 | 0 |
| billingAccount | 8 | 0 | 8 | 48 | 0 |
| catalog | 14 | 0 | 14 | 53 | 0 |
| category | 78 | 0 | 78 | 78 | 0 |
| individual | 36 | 0 | 36 | 72 | 0 |
| organization | 14 | 0 | 14 | 28 | 0 |
| product | 3 | 0 | 3 | 18 | 0 |
| productOffering | 39 | 0 | 39 | 39 | 0 |
| productOfferingPrice | 247 | 0 | 247 | 281 | 0 |
| productOrder | 15 | 0 | 15 | 105 | 0 |
| productSpecification | 25 | 0 | 25 | 49 | 0 |
| quote | 9 | 0 | 9 | 54 | 0 |
| resourceSpecification | 5 | 0 | 5 | 15 | 0 |
| serviceSpecification | 10 | 0 | 10 | 26 | 0 |
| usageSpecification | 6 | 0 | 6 | 36 | 0 |

## Error Summary

| Error Code | Count |
|-------------|-------|
| MISSING_PARTY_NAME | 28 |
| MISSING_PARTY_REFERRED_TYPE | 15 |
| MISSING_RELATED_PARTY | 293 |
| MISSING_REQUIRED_FIELD | 346 |
| MISSING_REQUIRED_ROLE | 247 |

## Detailed Validation Results

### appliedCustomerBillingRate Objects

#### Object: urn:ngsi-ld:applied-customer-billing-rate:7ea1db4a-b6b6-489f-ad68-1e5d0644dcb6

- **Type:** appliedCustomerBillingRate
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:applied-customer-billing-rate:a886304d-d699-4adf-b93e-dcdcd54474f1

- **Type:** appliedCustomerBillingRate
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:applied-customer-billing-rate:c25e45c6-8116-44d8-8101-651fddd379e2

- **Type:** appliedCustomerBillingRate
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:applied-customer-billing-rate:144d5a4a-6b0b-4308-beec-f15cddce3cbd

- **Type:** appliedCustomerBillingRate
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:applied-customer-billing-rate:0bd7ffd4-4a34-4e9c-8dc8-dfd60818bf65

- **Type:** appliedCustomerBillingRate
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:applied-customer-billing-rate:9b2925a3-775b-4996-8f7d-7d70ca15e367

- **Type:** appliedCustomerBillingRate
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

### billingAccount Objects

#### Object: urn:ngsi-ld:billing-account:3b49919b-ab08-4969-ab53-cfd06cc21206

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:billing-account:d2c224df-b007-4524-9029-7e7e1b021d35

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:billing-account:c84d03ff-fc74-435c-a54c-fed6e95ff80a

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:billing-account:f98a4654-612b-474a-875b-107243e814bb

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:billing-account:5a7a0a3f-61f8-4b62-bdf9-5212cbfc129c

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:billing-account:f727fb85-51f5-4dca-822f-4f4bb2775549

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:billing-account:4da08929-a8a0-4362-aed8-a4d0d25032f2

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:billing-account:b62afbb3-0991-4c3d-85be-d96f95c55435

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

### catalog Objects

#### Object: urn:ngsi-ld:catalog:1d535a9b-212c-4e8b-aaff-8a412e61dd0d

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:catalog:e40feb38-04b6-485b-9edc-789704d3cd85

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:catalog:03159bce-35ca-4938-bc2b-b8239e1008ca

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:catalog:c016bd48-8d84-48c9-82f1-8cb1d0c1cddd

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:catalog:4a86c5a6-610f-47b0-8c0f-6bbf472409da

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:catalog:44e13ee8-c1ee-4c7a-a0f0-a308cd73a0e3

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:catalog:ec45d9d1-50a1-445d-848e-09f50ca7862e

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:catalog:ad73d40f-11b6-4275-ab8a-6d1983e7b3ee

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:catalog:fb58a661-0341-4b61-9dac-610e228ad6bd

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:catalog:2df6d831-26d8-430f-9ac7-79cf447c47fe

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:catalog:721d9e67-0a46-4126-a12e-8f91670ceaf7

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:catalog:dfd8bf69-cb7e-4d0b-a3de-f407b2849580

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:catalog:b2a1b435-d234-4f58-a23f-92c6ba5798b5

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:catalog:134206dc-e658-4f12-a59b-ad2e17f4ede4

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

### category Objects

#### Object: urn:ngsi-ld:category:8d222bda-159e-4957-8b76-0fb06b4449dd

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:d6aa17c8-a0d4-4312-9ab1-ae78a0aace1d

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:73dd1ca0-defd-4642-bbec-f44d52273973

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:3867db98-1eb3-4bff-8a23-715082f404ff

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:d64a7aff-f49c-4cbf-84d8-4b883e59e392

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:ef149b50-08ea-4ee2-bf41-a90b04a1ce79

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:55b2cb9f-e13b-499e-bf76-fd1eb56d2bde

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:57158a09-c62b-4160-a49e-2295ca068682

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:3cad4d15-bbe6-44d6-a82a-a6478b8e53ce

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:35182b9c-be04-4a40-ae67-f9433f4d36e1

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:b332fe31-ec4d-4883-b051-e663a3d40830

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:3ae62a2d-a194-4250-9b3f-8c9349ef4799

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:bb19f802-c0dd-4686-9049-cbd4665e7a04

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:44370eaa-8828-4e1d-b9be-03b64eb11309

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:44e33c19-dd19-4ee9-a816-9dea2d34f02c

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:a7779dea-c0c0-47af-b078-3098c309cc23

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:840bf629-6e64-4e49-94c0-740d7b312dd4

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:7c796f0f-8e89-4578-bd6f-5d6c56ec912d

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:d154f150-2980-4b5d-9aef-2fdc123dc4d9

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:dc159e2d-d0a1-49e2-a2ec-b7c4c97e4dd7

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:2b015a9f-50e6-453f-a251-e8770e2260bb

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:7aa7ae47-4c3d-447e-848b-60b3f015cb49

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:9bfdb312-ad44-41c5-9c5a-4d58ad861f63

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:15f52622-9097-4fba-a66b-435fb651a393

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:ed12cd4a-3a23-46fe-8135-ea5bdfbb2b4b

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:bd99e5e4-5143-470e-9c05-d7839dfc71bc

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:b51bdeb5-0cf0-41a3-b0af-3ef48a88c988

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:17c3aafe-2c2d-49be-8dce-dd4981e22e28

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:ffeae64f-edc2-4754-906e-a4d35d10c806

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:91abfd48-4870-4af1-b64f-69cb144d1dc6

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:a14a4067-a5c9-4122-b2a7-28ee2dd85036

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:08892d05-2e7a-41f9-8e8b-2dcc28f86b8a

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:abe28b41-cc3a-4624-9c23-2ed3e021dbb5

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:96778468-41b5-4245-842e-f3c481a240b0

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:8337bace-8228-4925-a75a-e2654f0bac41

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:c053c314-6696-49bd-a10b-052e6f3dbddb

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:94bd8c25-4183-4124-a6ed-0516d4ced4a6

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:630359eb-7d66-4eb2-bf07-2590841049a0

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:af87fc59-aafb-4301-8ba9-3594558a9eff

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:9c02c10e-1cdf-4de3-a99e-f0b689881d2b

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:8795e633-8a64-438f-8cad-a2d5175cc9cb

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:47eeb384-4400-49c1-b0e5-718ef17860ba

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:9a7f4ac4-e481-4303-98fc-b4f9911f6529

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:03e76d62-a8c6-4571-ad7a-c31e754d3d3f

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:6ff41b67-bf97-4772-86ef-52d1d4998a3e

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:228fb947-48df-4281-b95d-c675bdc31b7a

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:419a8fac-d55a-4ae7-867f-2a3a338844cb

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:061ba18d-4263-42f3-9d6e-17d90ad25832

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:0b5cb3a5-b08e-4fca-8fb4-f99afaf71a2a

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:b1f8c6b7-e73e-415a-9f28-51c01adc36c1

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:55ff6485-3bcf-4009-815a-fca65727c64d

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:cfb68599-fde0-457e-bd8d-f6f6db20025f

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:13837a4a-4eee-4879-be8b-da8bc88f8086

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:10ba35dd-4906-4434-adac-9bb9268e1127

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:ef80f5f1-c480-48f8-b96b-b893efec3e4d

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:b4998d10-d9c2-46ad-a08d-63df7cd8e911

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:3beea627-9691-4ed8-9bbf-7c6fe06e278a

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:3ebbe833-9f85-42d2-b76d-7429d156eba6

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:818597c8-6de7-4961-8ce3-61fb30d7fbb8

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:a456f7a7-cbb0-4b71-9c25-f8cb46555016

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:2cc758fa-80b3-4704-9c38-26918c25b535

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:9c12beda-8e3b-4b90-ba96-191e0129a9d4

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:6df17fbc-465d-4e0e-93af-ecc0ccdf34b1

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:f0c23dff-8f09-474a-8306-607ca835d8ce

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:c9f5f9b8-2991-4cb5-aad9-eede449c88e8

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:4c9b41bb-2eeb-4229-8442-b0e8d8127271

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:917fa857-f354-49f1-84f3-38540e07e434

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:ac6eefc6-1e58-432f-a9b9-c8349c3a3d68

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:465a29be-b24c-448e-80ed-c9a1e5aeb3da

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:63ed0da1-bc76-4e8d-af4d-f551d8148394

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:1d731f47-e2ce-4f2a-899b-f7f1c449cd98

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:fcf925b0-bb66-4860-8c2a-db17861fe105

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:8ec9aa71-afa0-4ae4-93d2-61eef79fe9c6

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:2fffafeb-64ca-4d72-9e4c-ed7d2973b3a9

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:b664a377-e12c-465b-ae2c-0c5a53438e65

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:a6f7ec89-e2e3-429d-a16c-30b0a1ab0ff6

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:814efafb-b2be-4d7e-b83d-0d447d9f68ac

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:category:4b11ef5c-50a6-4a20-8817-8152b372cc95

- **Type:** category
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

### individual Objects

#### Object: urn:ngsi-ld:individual:97790bad-066f-41fa-bb5a-745d783e044e

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:038e6cf4-8805-4905-a151-42012ddcdfbf

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:c1046c38-844f-471f-873c-78b35c270c27

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:afa0bf6f-0b1e-4d9c-8529-e57bd28740f8

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:911f60f5-9a54-4d1c-ae99-46079d5c6143

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:9506e609-cf28-4302-9592-7a49521d0412

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:c5b57368-f3cc-46ed-838d-25153f471c5d

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:d6cccbbb-3fae-413f-89b9-c3d41ea25362

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:73ecc3c1-41b9-48a3-8c99-b29e6e02ad2d

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:8f307e88-4208-4484-9a62-b06b9b4d4a29

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:9e2e4859-c2f4-491c-a787-5f3aeec93ef6

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:b8d107f5-5888-413c-a334-df87468398ad

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:d3b2d0a8-d490-43df-8f00-f76b45cae5af

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:68f857d5-4de3-4f22-8daf-0b46d9699a8d

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:5c9d40b0-5628-4dec-9228-08fbee4bcce2

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:46b73e37-a453-4b19-8df0-85356a52cd8a

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:b266d7da-0852-4659-a003-820b391f7927

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:98d935c4-5f25-42de-a000-844950f1b766

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:0774f8c5-073e-4334-97b2-5b71167dfc2b

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:a0e4422b-1464-42a5-ac76-2901329a5c8b

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:b068358a-d4c0-4f20-a522-ac6286181d49

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:c475f14f-97b3-45ec-8a34-4f22f4931f04

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:f366c06b-38b2-4039-b66a-8dde9a88f1f2

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:80bb5420-a8c4-47a6-aa73-8c96ccc9a525

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:6aebe06a-54d2-4f17-87c4-940c1e22ef91

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:bb26a135-8281-430c-b2d0-652b752b0c15

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:5e7c0d3f-235e-4507-9875-a166fb41fb35

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:551f1882-63c9-44fb-b001-d9c2648ba6bc

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:16a87ab8-2d62-4298-97dd-6c96f2556397

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:521e32a5-bdc3-4530-a878-644299792594

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:2419897b-0da3-41f2-a34f-d92ccff3a9e2

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:67c4985f-2ed1-43d2-a631-6261a66f7c8e

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:6176fae2-ef4e-4bb5-aae6-f7d6e41efc3f

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:c6abe161-5c9a-4d56-9817-917162f1e453

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:b8be3518-b3ff-43a0-b8a4-5310a9ada50d

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:individual:b1256dde-1380-4ed5-95bb-8ebdcce08774

- **Type:** individual
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

### organization Objects

#### Object: urn:ngsi-ld:organization:00026ead-d16c-40ee-9218-48defc7ce749

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:a2f5ebea-49c9-4015-a9d6-56f2c566f6c9

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:d8dcc7a3-0774-4824-b552-d05d91986565

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:c2c39ab0-1f15-40f9-89ce-fc131953d33e

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:2bcbe859-e316-42f2-919c-f470cff9e235

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:eb6647da-84f2-4645-8d9f-c2905775b561

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:95fdc12e-6889-4f08-8ff8-296b10e8e781

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:e4fa0e9f-1779-49c0-9e8a-66a02bf1fe4e

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:4eb54a0c-a916-499a-bf5f-6bba76101e1e

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:fb47dd79-6497-4ec8-8456-e6e06d3c698b

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:44c33614-df98-459e-b710-064b1a7c6f65

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:6c53e937-212b-4e8b-997c-4d8695f789d1

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:338282a4-3c06-41d0-8c35-3fe5cecc38db

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

#### Object: urn:ngsi-ld:organization:3c9beb13-1187-4d67-9ecf-5e6825ffb2fd

- **Type:** organization
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)

### product Objects

#### Object: urn:ngsi-ld:product:45eaa84b-122f-4e6d-9797-31ab1ab16134

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product:b0a89e11-6fde-4bd0-948b-508088655e63

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product:34c6acaf-e895-4747-b371-6410d05e6a73

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty buyer.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

### productOffering Objects

#### Object: urn:ngsi-ld:product-offering:005974a1-f327-47bd-96fc-2c263f2818c4

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:4174966c-aca0-4787-88fe-96dd235fa2df

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:17f86010-d4c4-4120-99f5-25d25f376bbb

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:9bbc3d54-daae-414e-a63a-52fa06dfcd0f

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:dcdd91b1-22ee-41e2-968b-24289c077e18

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:7bbf2620-43fe-4afe-b2ff-cfe83f78484a

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:1c9bdbba-e2bf-4b24-8586-9d01f8900cf1

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:66b48359-1d47-42af-a974-1c40ee50f3dd

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:54a9af6f-f353-48a5-b599-84a535f0bc74

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:4699bdc1-592a-42b0-ad2d-93b0a06b044a

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:038df757-e986-4822-97a9-9934c903340e

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:2135b6c4-6de8-4ec9-9790-a3a322650a8a

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:fd2bd8d2-f504-4254-9d2b-f248534d4351

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:887a7ba9-2626-4f88-9252-5bc20239f61a

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:98f94a51-6c8b-4778-a39b-10f3428428b7

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:7e1e9a12-1cf1-4eda-bacf-803cb5d1ba58

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:f46853aa-44dc-4113-a3b6-bf4223b9e731

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:7bb70163-c7f3-417e-8566-d6691ba3d2d3

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:d0ea7910-50e9-411e-8015-a9fe4f6187f2

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:bdfbce5c-cbe9-426b-9c37-36696fcabe08

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:9b80e4ce-b813-43c3-ae13-689626c2fe67

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:da73b355-6f48-4f5b-8eef-606f4023366a

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:6aaf22ff-d5a4-4174-9a00-167a7d135b80

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:3d48f1c9-e4ac-456e-96eb-6268c49c6ed3

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:fd82e71e-1d49-4793-944e-5ad1dd261d9c

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:1c0b11e3-a36b-4851-95dc-65ef99cbdc26

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:b43daadf-db7d-48a2-b580-c9ee8ef6b847

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:220f7da7-945d-4b35-a319-77ddb84d5a7c

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:e212e631-2c84-4ca0-9226-3a3aeb9d6372

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:4111df70-5958-4ad2-a99b-1aaa56d0280c

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:ce4631ee-c8b6-4b81-8b65-229b71804de6

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:c0348e47-0554-40e0-bb93-02a38688dd7a

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:da30096f-def9-44ec-bde7-4f79048359c1

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:89993e7e-3c55-4679-abab-5d83ae51fb43

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:17e07571-2be1-4801-9f71-c715da95d515

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:9e246e8d-edc2-4843-b86b-3f461dc7b9fc

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:da74fc25-a1c6-4592-8e69-ece3262aa837

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:b3d9edd6-d09a-4752-a834-0bad9fc76b6a

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering:14642972-4178-4803-8485-aeb24b9df1a8

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

### productOfferingPrice Objects

#### Object: urn:ngsi-ld:product-offering-price:91a5b7f3-afb1-427c-bed1-85332ee1448d

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:0253d786-a05f-46b2-94c7-c5e5654b330f

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:5a10aced-486c-461f-a0b7-bc9e60fd9c13

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1bb33bf5-d924-4155-8cbd-0771aff9bc1d

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:03d505ae-b938-4449-9de2-b766127c189e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:a21738bf-5fe3-4f09-8f86-7b57bf339403

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c1a2e7d0-cc93-46c9-a6c9-26e68072a11f

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:7b6340ad-0087-4abd-9342-ba29f90dc3bf

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:ad27d2ad-f48e-4c2e-826d-3106a75fd7b4

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:33e2752d-d9d8-4bb4-afa1-c63238db3bd7

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:3fb637a3-1c25-42e5-8cb3-d4cd284ab302

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:dfd79425-ca90-4134-9a39-6ed4c2990eb1

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:74bceffb-1389-4b71-b665-5e67d2d6f07c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1094dbc9-1816-4fba-af99-4e1581042336

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:56980c99-2847-45d7-baf4-0b9a42937615

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:05c0cffa-a599-4498-9bfb-e5852f960b6e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:56a31088-be75-4da2-8fc7-4dc267179cb5

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:6315e516-9989-474f-94fd-1c27db780920

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:a006b275-a16e-45a3-ae01-fcd3c04837af

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:2425cab1-5fc4-4557-b964-500dc62efabf

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:584aa66e-2508-48aa-9fe9-22014e27564a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:a3e5298d-07d1-4b97-b9cc-39c723afbd97

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:425925c2-1494-4767-997c-d2dc5939780a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:def636a1-8110-4779-b131-4a22401cfcc8

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:7a4802b5-ba89-42c8-bdc0-7bb70977469f

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:78b0873b-d229-4d50-bfc6-540a5aa7839f

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:545959f6-f3d7-444b-bd90-82a19ef8183a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:cf257bb4-3c7b-45e5-a9a0-fece3371ee7e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:7fcfb599-157c-4ecd-a200-d7adbb2abe9e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:f6b0eaa7-08dc-4338-af27-dbf0afe764b2

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c807d97b-04f2-48f7-905c-dbc0432eea55

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1bcf1d19-59ed-41ab-9611-6c73de8804ae

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:109c4c00-9463-42c1-b8f2-1388f9317992

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:001cccc0-3dc9-4785-97f2-8ab60b19f194

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:be83fb10-fe4a-4f99-aeb0-212a07ac0ec5

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:955a60dc-6caf-486b-a037-906bf2206b2b

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:ddaf6425-04f5-4735-bf65-eec711cf3998

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:3d267f84-cae1-48ca-a296-ec1c93b77183

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:aa7e2451-b7c0-46d3-9bec-62afdb0bb051

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c08efc8d-7e06-47b6-8521-4f1b48890f05

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:911a85ed-e581-4aed-91c8-7074a2ffec84

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:f157384e-16d0-462b-a499-46cb37a2e2d0

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:0c88f904-2b70-4e27-bfdc-0ab019831dba

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c06b2f83-825b-4dde-96c2-c3152c2233a3

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:bb12691e-0307-4938-b610-7455d5c57bed

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:97e5f94a-95d9-40f6-af96-310a05598502

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:fa954dbb-b264-47d1-abce-7d48a8bddc19

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:7196cfd8-e2d4-4880-b97b-f9639641365a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:d13378a7-bc3a-4cb7-ab81-35a632ed0c4d

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:727971fe-dc68-4cf3-8252-7d1f22494e09

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:55a279de-c4da-451b-9f04-00571610cf69

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:47b8056f-4ca3-4c4e-b24c-dcf9edaf7c44

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:f2247dcd-d3fc-483c-9b55-f2d7605d952e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:eaae2846-1cc6-4410-b452-d82cc083dd75

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:fadd7fa8-e444-4879-bf6f-5808109b69e6

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:183bf946-4173-469e-a1e5-0fe9a596cc5a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:19dc89fe-1d3c-454f-ad1f-605b8190409f

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:b1be3749-046a-4ad7-8265-82cda63a138d

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1af64f8b-7756-4bb3-8b78-eab759274f77

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:e16b0922-907e-406a-9255-62f4b49b3be1

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:bdaacde6-3fcb-4137-98f0-a3964c35cf37

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1ebccd8b-e227-47cb-b8b3-ff5c7e4a1462

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:8c9113fe-05d5-4655-9d4a-ccb5509bf0fb

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:09d2a6bf-d15b-4e42-a85d-b713c8b76670

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:cf62c14e-4304-46f9-92ce-13784d3852b7

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:2e956de4-80a4-4ddd-a956-235ea9695b9d

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c6113c55-35b2-4e2e-9c34-94f749aa514b

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1bb7a278-1fe6-4ecd-9b9d-cbb2cc8c0a3a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:385d401f-6b65-4fd2-b9db-7f8552dc6296

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:7351a8bb-076d-41a3-af04-25898a5c093a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:e1032cdc-b18d-401b-a5f3-ad1e3a2d43f9

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:f9ff1dab-392c-45d4-9ebc-00f01641d6ae

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:4de020d9-2c51-4c17-813d-8b78ecf93a19

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:40407786-84e8-400b-86b6-854a3aad8ec0

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:ad03f5d5-2f08-46d2-82a5-c754643ad7e8

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c98700d9-2ec4-4ef2-8edb-162280e47eae

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1fb01b3a-7f98-4091-9be4-b51ce7523a65

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:69b4645e-6827-4941-9bb4-dfbcb27e6a29

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:ec7b8370-5ce3-4d54-9622-5338590832cd

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:0186370d-e049-4632-ae04-38d5eca8816c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:a199e040-9a6b-4a1f-a755-3c43852739e2

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:60c80e01-4f39-43f6-88cc-ff8ea674f87a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:47983a37-a4df-4d04-b005-5e2f8dca7fab

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:ad91f4de-b402-40f4-a6af-e5956e9f00ba

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:56e78737-ada5-47f7-8902-4aac351ed3a4

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:249e1126-b160-4680-8d87-6e89b812a817

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:9b4406a2-df4c-4ae8-a3cf-a6246575641a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:90543597-4b18-4012-8c6e-9f26109f42a0

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:e756395a-4b48-4df4-9f98-d1624d3e4212

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:f3e66a3f-b733-432d-b22a-ea9dca50c48c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:54fe7d22-8fd2-4fe4-a9dc-8e27aab98bac

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:cc4115c7-571d-4e0a-b752-199850808527

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:71e7fb14-7b06-4bc0-98a9-8a26f0d72117

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:2c4c2460-7e18-47e7-ac86-3361f66eed06

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:ff1bee47-15f2-4800-aa6a-3e8b4896dd0c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:901f2b9d-836e-4fb7-929d-6970b581bed4

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1662f498-043b-4a6b-8dd5-c5aabea025a5

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:5479bb4f-7ade-42b8-828d-d8685780de94

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:a90ba1d5-3e77-4852-bce4-dd2478b46763

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:3d2f95df-c6c1-4794-88ff-b652580ef119

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:7cf80cee-8e0f-49e4-8960-1d8fc1cc9b6c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:74697e82-ac10-4d0b-b270-1bde5bf70f4d

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:0a90a84e-9d09-4b97-a7dc-7f5a69afc6e0

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:a71ecbb6-01e7-48d1-8bf6-0ccbced7e888

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:bc31286e-324e-499c-9db8-df759cee12eb

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c27f73e0-54de-4def-bf1e-2b418b4df5c3

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c0852f36-0dce-4193-b823-6580756fb8cd

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:94876797-0459-4288-8628-bdad9c9db907

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:4cc26f3c-a33c-4be1-9f44-55283548f1ce

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:be8f6202-e4e9-4d32-bfaa-bc9db4efaa41

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:e7268246-9ce1-4392-8195-5438573faef7

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:cf2aefbe-9e22-499e-9725-4a0834f7acf2

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:b521af25-53c0-4b5a-af88-8cbae6d54d3a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:7ea1cb0c-1dfe-4c8a-90d0-64d2605fb4ee

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:2b153541-c97b-4449-9d9e-0ef553016cca

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:fa450e4c-693d-4db6-8523-b7176bcd2c54

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:34c01a67-042c-4815-b1a2-331f7a7de2e0

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:7f5b3871-db5d-4e37-acaf-699ebe423bf4

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:07485fea-ab68-4887-b21b-ed38a428d8bf

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:87a0d1c6-976a-4c99-8272-d464bb740ee8

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:5d027b1f-3975-4f21-8436-42df8ab25cd4

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:49f2f357-db82-4d6e-8750-7a39b1d3ad77

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:f6a70f23-4509-453a-af0d-58da0bdaee1b

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:4896d719-c5ff-4032-bdcf-8921b55afd01

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:4ee4de99-0d21-4023-b97f-07d9bce8649a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1c0dde70-1a2f-4ffc-9a08-95e00f4d1204

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:4270ba22-aefc-475a-bed6-66964a24a1ba

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:71a69afb-ea75-43f7-b52b-b1e0b7eee845

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:f4a86158-f6df-4636-b3ac-e888ad28fc02

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:49e642b2-930f-4b50-8ddc-dcd17026afee

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:dbc9e96e-0317-4dc0-b766-57e827e9e637

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:cf0d9799-d9ff-473e-863c-5e9b980695b8

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:145ac339-3153-415b-8b24-176ec0e1cb4f

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:48a0f325-78d1-46be-ad39-5c69be878ecd

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:984060ea-7556-4f33-bb8e-26f56efc69d3

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c86d9a42-1365-4913-9e6c-bf39789eba9e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:d859418f-8464-4b8c-8260-3a52f8be3033

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:8ce019f5-fabc-4dc8-8498-cbc617aac227

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:bcdf7806-6022-4bb0-8cd8-0ef4b3ce495c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:4504988d-fd5c-4e8b-a8e6-033356157e95

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:8ef944ef-2ef9-4fdb-8485-c34dd5fe94e3

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:baff3d62-a51b-4ca3-b0fc-2958356a1e72

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:5b14958e-fc2c-44c8-8029-43b12a91273c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:d60e5d5c-b83d-4ab7-a706-6d1b34914dbb

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:6c8e986b-8deb-40c4-9141-cb8f232c377d

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:43155aaf-52c5-464a-af9a-fd50ba1f6851

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:6b34081a-1b1e-4e06-9a95-b95aca58370b

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1d9fa601-08c5-4b0d-b0bc-63c33a03e0a0

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:02b6859d-0bde-433d-b03c-57a742b3a1e2

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1b7e22ec-00b1-4705-8393-fafb88f76524

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:96b8ca9e-dee6-4238-83c5-d83ba1fcc4a1

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:d64e6b39-1f6d-4a5b-9c1d-c8f77935f2f6

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1e9ba62a-0156-45d3-a6b1-854fe6899de0

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:96ff38f9-765c-46d5-935a-1076bfbf9833

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:acbd47c3-ea82-426b-af67-c7e8aafde298

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:57113c08-48f5-4c7b-a073-fb61133ea394

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:35e91a35-dbe5-471d-b090-4ddf0ea3db2a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:ff4ab74d-1e82-4929-b292-7b08366e161c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c17b48b6-fa66-4b7d-aef7-230d8e437b4e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:0d0477d0-48a4-4d32-a42a-f2baef1691fa

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:8532d238-5caf-4503-a87a-751a29a53b81

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:f465bf64-d221-4fa2-9c56-053b7f8a343c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:330d4219-23d6-4533-ae84-2be3cefed8ec

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1c329beb-e6b1-40a8-b15e-11c37e5e69b6

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:4b653124-c470-474e-99b3-03282914f162

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:0442a0df-e144-48ea-8c58-42cf414bc4b7

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:3a2682ee-b975-4bfa-9a14-4bd60d00162b

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1e3324d7-bf3d-408c-9f16-543377e0fb1d

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:578e28f8-67da-4369-a06c-3cba0653dc57

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:b5ac3f40-97af-4370-8a88-c7659e119af7

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:f4ce38ac-8dfe-459a-b4fe-c3d5305a8593

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:53d7abb1-78aa-4631-84e5-11563bbb00dd

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:ec528142-8ed4-4ff7-845f-4ea1822a5723

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:cbcbdbbc-9dcc-4976-85d8-2c9d7e210cc4

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:450d93f0-3439-4b35-b09f-c8b0f965cc6f

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:edd68732-afc1-4cac-80c7-550f4b531a5f

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:9981eb78-2629-4b89-bf70-55b2cf8c370d

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:54263e1a-a8e4-48e3-b6cb-7d5f89be0cfc

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:d6ef067f-2c64-4557-bb13-cdd6dfd6c5f7

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:f49ad34d-397a-4a8b-b2d1-a6534b7bb43b

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1c8d89c9-2609-49dd-b760-73f5cffbf97d

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:9e9f7c77-5168-4174-932f-48e9cd6918db

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:e5a8a3c4-dff3-495f-bb08-c095cf04e28e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:d4de9ad2-89ac-4325-96e9-26aabc49415e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c1d6eb86-8fa2-4629-b545-ba58d3cc0098

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:3ad70bb3-9545-4c81-bb67-8872a56b6833

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:1d8044d3-3281-43f9-842b-cb0974a0881a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:a6d71be1-57e1-4555-bfdd-f483a0c637c7

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:61a410c7-99d9-4da6-84a7-1e6ef98350eb

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:781c55d3-77b7-4d59-bedb-6a186524bbb6

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:89f930b4-5601-44de-97dc-dcb173dbf66e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:d7c32b40-ad7e-4019-865e-b655ae80cad6

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:deace989-a7ba-4e0a-aa8e-fbf81a954f13

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:9528e528-4b97-4520-934a-d2d31c545b35

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:2d47fe3c-64b1-4f47-abc8-46cd4cf16dba

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:011a3164-3193-44d7-8b70-8945ddd9ae88

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:54d21df4-f997-45cd-b7ba-f749974a409b

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:45b80f4a-b4a9-4b24-b41e-319ece18995e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:92b0f334-a171-4c72-b248-d64eae9a87da

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:00be1ba5-0b0d-4204-b903-9111d9f854d5

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:88903eb8-9093-484d-bfda-a6367bb0f73c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:7cfdb910-bc8b-46c6-8dde-005f38bcfbbd

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:8130520e-eba6-4085-9500-529d6d3c7b69

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:bbda5de2-fe37-40f4-af5e-bfc3cd75bd3e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:7b61a8d4-bcf4-42c4-a7b5-531661662532

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:06c2feb7-fd7f-464c-b692-c6e77489826c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:d7f472cd-99e1-419e-800a-ae3f1d735ba0

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:0d3b58bd-f011-4f7e-83df-ab207f4182f9

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:4f25a73f-81b7-43c3-b7fc-28535da9ce32

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:79e80832-752e-4717-a641-b960447cfe39

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:4c7e5965-cb28-4c1b-89de-35d9da8b69c6

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:73566286-27d4-4739-9458-f54a0d9acf2f

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:21d14112-4193-4e90-bb3d-225839531cee

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:690978ac-a697-4f2d-ba20-c971a56ff62a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:aa741a8b-bbd7-4879-9f72-76be7abb7993

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:2127e6c0-ec68-4f81-8b7e-ce3863214f9c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:ab959354-76d9-425a-a47d-27d6210300c0

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:a98f47e2-0fab-4e5c-92d7-413a4fd83629

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:fe269e90-dd37-423e-9433-6d2888c78530

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:86dfbbc5-e561-4354-9043-25b353e96a94

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c74e7f33-6812-4252-89aa-4fc8f9cf53cd

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:923341bb-c819-426a-a07a-de63c75526b6

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:66f34e25-5b92-464c-8b67-f166e4633e6e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:f135cc3d-f8b5-4204-a162-9c30241d3a7a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:af93998a-9d33-4c7d-a4dd-50ad21fce26f

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:bcae3d57-6ad7-49b4-9d33-8f21a61c305e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:8cf96f01-69ee-4809-893e-05655a8e24ad

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:b54cb8d4-a9c9-4fa9-a683-450b5f547e4b

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:c65269e8-41a1-477d-964b-52dfd3cb215c

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:28f6729d-e998-45d1-be16-28d5882373f6

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:8ab6efb5-c792-4746-9ae8-4b11021adca3

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:125e4594-1ecf-4f4b-a4d7-1707a4750e86

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:4f0fb6b7-3b6c-4841-ba34-de3ba9ed6059

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:b2d7856b-eeb2-4537-89b5-33fea3b307ff

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:b31c770d-7038-4340-8f02-3cf78db7ad0e

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:5e3fef83-5502-493e-b714-b32d8ae28d9f

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:16b3904b-3e9e-4d57-a276-189ed50bf7f2

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:837493a5-659d-4a6a-a31d-666fe61549c2

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:75b407ef-8cbb-4130-9c27-0500e7786684

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:acb4594d-0e2c-4a2e-a314-a420801fa1a0

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:29183343-4c38-43c5-8b4d-bdda66fa069a

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:23db4cb5-ba4f-4306-9966-112073445de9

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:fe4f0368-e4b9-434e-b86f-06ea0757cad1

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:e36dc635-2d6e-4c3f-b208-10e303d287f3

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:948b22bb-9b16-46a2-9d3f-954fdbab57a1

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:56cb49ad-9694-4b46-901f-ae70ab6e4740

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

#### Object: urn:ngsi-ld:product-offering-price:07345458-dfac-429e-9610-d4ed40ede667

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

### productOrder Objects

#### Object: urn:ngsi-ld:product-order:fdc85cc5-6e74-4e67-a9d0-1f3d7a79dce3

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:4902aa01-8949-4a46-9be4-1cf98aa4a4f7

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:117780df-ea0f-45b9-bcb3-1cbe5b17b1f9

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:06350c8a-4238-4284-a9b4-f5c047b91292

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:9dee5ca2-947e-4443-aca8-f82ec936d2b2

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:fa5e91ba-6418-4354-b840-1f9e31ca1e74

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:abee6b6d-f3e0-4465-bdac-bea275ddecc2

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:1c4401c3-4770-403e-af95-6b9737bc087a

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:525f2198-45a7-44a9-a0d3-413972bb5f56

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:bfdd1be4-a32b-4017-855d-d4fdaa95818a

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:0c6cc039-21bd-462b-9a45-b6ec77245e96

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:367cc7f9-2ea7-4fc8-b9d4-1f2aff2e57d0

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:4ffa78f5-2b85-4c9e-b09f-00d0ecad3a44

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:d7971e80-dd40-42c3-b33d-1a5485dc38a4

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-order:18777b58-f7d2-4570-ab6d-2917f86b3c55

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:32 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty seller.referredType: Related party referred type is missing (Code: MISSING_PARTY_REFERRED_TYPE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

### productSpecification Objects

#### Object: urn:ngsi-ld:product-specification:bf72e349-03f0-4e88-965a-73815d8881b4

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:55158c3d-ef8f-4f1c-b9d0-82dd3138b2ae

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:5e2ced54-45f1-4687-b9f9-ee13cb318a66

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:ccaec925-9fad-40c0-8015-5b9087f92aff

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:ba4123f5-77f2-4b80-8c8d-13c01d3a6e72

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:411816fd-3f8b-4e6f-8c35-02e8cada5ce9

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:534697ce-c774-4cc2-a8f4-d4f2ae6cf4f9

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:bfe7418d-2614-4e6a-920c-6d6711171f10

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:10ebf2c1-fe79-4191-81a3-b58207307c5a

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:2532e67f-0a84-4aed-b31f-7654984daac9

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:351bbc6e-515f-4159-9082-e60563c927da

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:697c9eae-0646-454f-a210-e20072941f7b

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:90496a79-8eb1-425f-b47c-314b15a1c256

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:905e3f69-423c-4d1d-a713-13650e1ec5a4

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:468c4c03-d5a5-42ec-9dab-3c099095e621

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:f57c1be7-094b-4051-aa1e-ad40c169ff00

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:f40073af-0cf2-4531-8a86-0d7928f0bf32

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:056e1bd6-f9ab-4b56-a651-e67f9e5c1545

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:9c5f5aa8-b400-42b7-aa4c-c98dcfb64df8

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:1c79f231-ab2f-4183-89c9-89a07883a964

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:e285d9e5-7b4a-49a5-95d1-4ec4b99eb907

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:532710db-3418-45e8-940d-1bbd25a2c398

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:653468a2-6c44-4877-a89b-e1404648d083

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:f25c9a61-691b-4e6d-abc3-34d9105d3877

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:product-specification:dd837633-e03e-4ca0-9825-f363b63b8766

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Related party information is required but missing (Code: MISSING_RELATED_PARTY)

### quote Objects

#### Object: urn:ngsi-ld:quote:91333dfb-bf4c-442c-8072-26ea2940b7ec

- **Type:** quote
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:quote:745edf57-9c8a-488d-8059-a2616bd743b7

- **Type:** quote
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:quote:294c4f48-e350-453d-952c-9132ec06bd0a

- **Type:** quote
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:quote:410e6d9c-da52-4894-bef5-5acb546a4bc2

- **Type:** quote
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:quote:014f8da2-2a39-4c99-a5d2-e22ce06de8bf

- **Type:** quote
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:quote:b1d1e81c-d152-4a4a-8e97-8f92dd502d9b

- **Type:** quote
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:quote:ebc5fd0b-ec81-40ca-acf8-5b39fe79c01c

- **Type:** quote
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:quote:ae0219c0-a236-4b6d-9cda-892c7ddda824

- **Type:** quote
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:quote:b8401b74-d7c7-4797-83d5-05e66d9a4419

- **Type:** quote
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty seller.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

### resourceSpecification Objects

#### Object: urn:ngsi-ld:resource-specification:9bd04c66-97a5-489a-ae93-7dddb4cd341b

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:resource-specification:f0153609-77d6-4464-9352-679f8d0e015f

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:resource-specification:e84e77c9-55e1-4c2f-953a-09dd52003f92

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:resource-specification:59cf3608-2879-4f38-a8f7-5baf61653c92

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:resource-specification:46e91850-687d-4454-a53a-b2211bc8f4b8

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

### serviceSpecification Objects

#### Object: urn:ngsi-ld:service-specification:ab1f9684-e04f-4692-bb2d-20827f1bb759

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:service-specification:8f5f2d0c-9af4-47ad-a932-387455fc11df

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:service-specification:69a374ff-97c5-417b-8b31-2bd36798006b

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:service-specification:80b2ba93-5d5d-4753-a29f-b80114a01333

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:service-specification:5ff52024-ea17-4f17-aa39-0499f57fc7d1

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:service-specification:8a8741bb-c559-46a8-9ae8-d2f00c7504ca

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:service-specification:89d6bddb-7805-4042-a60a-030ef09ed816

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:service-specification:f70f5f90-14d2-4442-a459-dae184f233e4

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:service-specification:d612a81a-b154-4ca4-8ed0-28c13233e268

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:service-specification:b4de64b4-d9cc-475a-8351-bb7b0a55a9ce

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:30 UTC
- **Errors:**
  - relatedParty selleroperator.name: Related party name is missing (Code: MISSING_PARTY_NAME)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)

### usageSpecification Objects

#### Object: urn:ngsi-ld:usageSpecification:21333496-4652-4bbf-be85-a897278d4ee9

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:usageSpecification:24abecd5-bf1f-42e0-a34f-1657f39dffe1

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:usageSpecification:156985fe-69c9-4f2f-a906-cbdc01d4d427

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:usageSpecification:3aacf05a-838b-4160-a810-4447fc58695e

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:usageSpecification:fd5da80e-fb29-45e8-8636-46b46dc2973a

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

#### Object: urn:ngsi-ld:usageSpecification:e16ff3ad-0bf6-4b8c-95d9-0d2e69b4c12a

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-09-11 05:42:31 UTC
- **Errors:**
  - lastUpdate: Required field 'lastUpdate' is missing (Code: MISSING_REQUIRED_FIELD)
  - version: Required field 'version' is missing (Code: MISSING_REQUIRED_FIELD)
  - relatedParty: Required related party role 'seller' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'selleroperator' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyer' is missing (Code: MISSING_REQUIRED_ROLE)
  - relatedParty: Required related party role 'buyeroperator' is missing (Code: MISSING_REQUIRED_ROLE)

---

*Report generated by TMForum Proxy Validator*
*Generated at: 2025-09-11 05:42:32 UTC*
