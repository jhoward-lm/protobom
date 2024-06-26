package unserializers

import (
	"fmt"
	"testing"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/spdx/tools-golang/spdx"
	"github.com/stretchr/testify/require"
)

func TestSPDXExtRefToProtobomEnum(t *testing.T) {
	s23 := NewSPDX23()
	for _, tc := range []struct {
		sut          spdx.PackageExternalReference
		isIdentifier bool
		expected     sbom.ExternalReference_ExternalReferenceType
		shouldErr    bool
	}{
		// Category: package manager
		{
			spdx.PackageExternalReference{Category: spdx.CategoryPackageManager, RefType: spdx.PackageManagerBower},
			false, sbom.ExternalReference_BOWER, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategoryPackageManager, RefType: spdx.PackageManagerMavenCentral},
			false, sbom.ExternalReference_MAVEN_CENTRAL, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategoryPackageManager, RefType: spdx.PackageManagerNpm},
			false, sbom.ExternalReference_NPM, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategoryPackageManager, RefType: spdx.PackageManagerNuGet},
			false, sbom.ExternalReference_NUGET, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategoryPackageManager, RefType: spdx.PackageManagerPURL},
			true, -1, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategoryPackageManager, RefType: "chacha"},
			false, 0, true,
		},
		// Category: Security
		{
			spdx.PackageExternalReference{Category: spdx.CategorySecurity, RefType: spdx.SecurityAdvisory},
			false, sbom.ExternalReference_SECURITY_ADVISORY, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategorySecurity, RefType: spdx.SecurityFix},
			false, sbom.ExternalReference_SECURITY_FIX, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategorySecurity, RefType: spdx.SecuritySwid},
			false, sbom.ExternalReference_SECURITY_SWID, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategorySecurity, RefType: spdx.SecurityUrl},
			false, sbom.ExternalReference_SECURITY_OTHER, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategorySecurity, RefType: spdx.SecurityCPE22Type},
			true, -1, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategorySecurity, RefType: spdx.SecurityCPE23Type},
			true, -1, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategorySecurity, RefType: "chacha"},
			false, 0, true,
		},
		// Category persistent id
		{
			spdx.PackageExternalReference{Category: spdx.CategoryPersistentId, RefType: spdx.TypePersistentIdGitoid},
			true, -1, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategoryPersistentId, RefType: spdx.TypePersistentIdSwh},
			true, -1, false,
		},
		{
			spdx.PackageExternalReference{Category: spdx.CategoryPersistentId, RefType: "chacha"},
			false, 0, true,
		},
		// TODO(complete this)
	} {
		tc := tc
		t.Run(fmt.Sprintf("%s/%s", tc.sut.Category, tc.sut.RefType), func(t *testing.T) {
			res, isIdentifier, err := s23.extRefToProtobomEnum(&tc.sut)
			if tc.shouldErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tc.isIdentifier, isIdentifier)
			require.Equal(t, tc.expected, res)
		})
	}
}

func TestExtRefTypeToIdentifierType(t *testing.T) {
	s23 := NewSPDX23()
	for _, tc := range []struct {
		sut      string
		expected sbom.SoftwareIdentifierType
	}{
		{spdx.PackageManagerPURL, sbom.SoftwareIdentifierType_PURL},
		{spdx.SecurityCPE23Type, sbom.SoftwareIdentifierType_CPE23},
		{spdx.SecurityCPE22Type, sbom.SoftwareIdentifierType_CPE22},
		{spdx.TypePersistentIdGitoid, sbom.SoftwareIdentifierType_GITOID},
		{"", sbom.SoftwareIdentifierType_UNKNOWN_IDENTIFIER_TYPE},
	} {
		identifier := s23.extRefTypeToIdentifierType(tc.sut)
		require.Equal(t, tc.expected, identifier)
	}
}
