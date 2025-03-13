#!/bin/bash

abigen --abi=modules/IdentRegistryStorage/IdentityRegistryStorage.json --pkg=identregistrystorage --out=internal/modules/identity/identity_registry_storage.go

   abigen --abi=IdentityRegistry/IdentityRegistry.json --pkg=identityregistry --out=IdentityRegistry/identity_registry.go

      abigen --abi=ModularCompliance/ModularCompliance.json --pkg=modularcompliance --out=ModularCompliance/modular_compliance.go

         abigen --abi=compliance/FinancialRWA/FinancialCompliance.json --pkg=financialcompliance --out=compliance/FinancialRWA/financial_compliance.go

            abigen --abi=compliance/RegistryMD/RegistryMDCompliance.json --pkg=registrymdcompliance --out=compliance/RegistryMD/registry_md_compliance.go

               abigen --abi=Token/Token.json --pkg=token --out=Token/token.go
 