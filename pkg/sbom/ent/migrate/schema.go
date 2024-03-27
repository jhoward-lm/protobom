// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DocumentsColumns holds the columns for the "documents" table.
	DocumentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// DocumentsTable holds the schema information for the "documents" table.
	DocumentsTable = &schema.Table{
		Name:       "documents",
		Columns:    DocumentsColumns,
		PrimaryKey: []*schema.Column{DocumentsColumns[0]},
	}
	// DocumentTypesColumns holds the columns for the "document_types" table.
	DocumentTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Nullable: true, Enums: []string{"OTHER", "DESIGN", "SOURCE", "BUILD", "ANALYZED", "DEPLOYED", "RUNTIME", "DISCOVERY", "DECOMISSION"}},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "metadata_document_types", Type: field.TypeString, Nullable: true},
	}
	// DocumentTypesTable holds the schema information for the "document_types" table.
	DocumentTypesTable = &schema.Table{
		Name:       "document_types",
		Columns:    DocumentTypesColumns,
		PrimaryKey: []*schema.Column{DocumentTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "document_types_metadata_documentTypes",
				Columns:    []*schema.Column{DocumentTypesColumns[4]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EdgesColumns holds the columns for the "edges" table.
	EdgesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "amends", "ancestor", "buildDependency", "buildTool", "contains", "contained_by", "copy", "dataFile", "dependencyManifest", "dependsOn", "dependencyOf", "descendant", "describes", "describedBy", "devDependency", "devTool", "distributionArtifact", "documentation", "dynamicLink", "example", "expandedFromArchive", "fileAdded", "fileDeleted", "fileModified", "generates", "generatedFrom", "metafile", "optionalComponent", "optionalDependency", "other", "packages", "patch", "prerequisite", "prerequisiteFor", "providedDependency", "requirementFor", "runtimeDependency", "specificationFor", "staticLink", "test", "testCase", "testDependency", "testTool", "variant"}},
		{Name: "from", Type: field.TypeString},
		{Name: "to", Type: field.TypeString},
		{Name: "node_list_edges", Type: field.TypeInt},
	}
	// EdgesTable holds the schema information for the "edges" table.
	EdgesTable = &schema.Table{
		Name:       "edges",
		Columns:    EdgesColumns,
		PrimaryKey: []*schema.Column{EdgesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "edges_node_lists_edges",
				Columns:    []*schema.Column{EdgesColumns[4]},
				RefColumns: []*schema.Column{NodeListsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ExternalReferencesColumns holds the columns for the "external_references" table.
	ExternalReferencesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString},
		{Name: "comment", Type: field.TypeString},
		{Name: "authority", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "ATTESTATION", "BINARY", "BOM", "BOWER", "BUILD_META", "BUILD_SYSTEM", "CERTIFICATION_REPORT", "CHAT", "CODIFIED_INFRASTRUCTURE", "COMPONENT_ANALYSIS_REPORT", "CONFIGURATION", "DISTRIBUTION_INTAKE", "DOCUMENTATION", "DOWNLOAD", "DYNAMIC_ANALYSIS_REPORT", "EOL_NOTICE", "EVIDENCE", "EXPORT_CONTROL_ASSESSMENT", "FORMULATION", "FUNDING", "ISSUE_TRACKER", "LICENSE", "LOG", "MAILING_LIST", "MATURITY_REPORT", "MAVEN_CENTRAL", "METRICS", "MODEL_CARD", "NPM", "NUGET", "OTHER", "POAM", "PRIVACY_ASSESSMENT", "PRODUCT_METADATA", "PURCHASE_ORDER", "QUALITY_ASSESSMENT_REPORT", "QUALITY_METRICS", "RELEASE_HISTORY", "RELEASE_NOTES", "RISK_ASSESSMENT", "RUNTIME_ANALYSIS_REPORT", "SECURE_SOFTWARE_ATTESTATION", "SECURITY_ADVERSARY_MODEL", "SECURITY_ADVISORY", "SECURITY_CONTACT", "SECURITY_FIX", "SECURITY_OTHER", "SECURITY_PENTEST_REPORT", "SECURITY_POLICY", "SECURITY_SWID", "SECURITY_THREAT_MODEL", "SOCIAL", "SOURCE_ARTIFACT", "STATIC_ANALYSIS_REPORT", "SUPPORT", "VCS", "VULNERABILITY_ASSERTION", "VULNERABILITY_DISCLOSURE_REPORT", "VULNERABILITY_EXPLOITABILITY_ASSESSMENT", "WEBSITE"}},
		{Name: "node_external_references", Type: field.TypeString},
	}
	// ExternalReferencesTable holds the schema information for the "external_references" table.
	ExternalReferencesTable = &schema.Table{
		Name:       "external_references",
		Columns:    ExternalReferencesColumns,
		PrimaryKey: []*schema.Column{ExternalReferencesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "external_references_nodes_external_references",
				Columns:    []*schema.Column{ExternalReferencesColumns[5]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// HashesEntriesColumns holds the columns for the "hashes_entries" table.
	HashesEntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hash_algorithm_type", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "MD5", "SHA1", "SHA256", "SHA384", "SHA512", "SHA3_256", "SHA3_384", "SHA3_512", "BLAKE2B_256", "BLAKE2B_384", "BLAKE2B_512", "BLAKE3", "MD2", "ADLER32", "MD4", "MD6", "SHA224"}},
		{Name: "hash_data", Type: field.TypeString},
	}
	// HashesEntriesTable holds the schema information for the "hashes_entries" table.
	HashesEntriesTable = &schema.Table{
		Name:       "hashes_entries",
		Columns:    HashesEntriesColumns,
		PrimaryKey: []*schema.Column{HashesEntriesColumns[0]},
	}
	// IdentifiersEntriesColumns holds the columns for the "identifiers_entries" table.
	IdentifiersEntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "software_identifier_type", Type: field.TypeEnum, Enums: []string{"UNKNOWN_IDENTIFIER_TYPE", "PURL", "CPE22", "CPE23", "GITOID"}},
		{Name: "software_identifier_value", Type: field.TypeString},
	}
	// IdentifiersEntriesTable holds the schema information for the "identifiers_entries" table.
	IdentifiersEntriesTable = &schema.Table{
		Name:       "identifiers_entries",
		Columns:    IdentifiersEntriesColumns,
		PrimaryKey: []*schema.Column{IdentifiersEntriesColumns[0]},
	}
	// MetadataColumns holds the columns for the "metadata" table.
	MetadataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "version", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "comment", Type: field.TypeString},
		{Name: "document_metadata", Type: field.TypeInt, Unique: true},
	}
	// MetadataTable holds the schema information for the "metadata" table.
	MetadataTable = &schema.Table{
		Name:       "metadata",
		Columns:    MetadataColumns,
		PrimaryKey: []*schema.Column{MetadataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "metadata_documents_metadata",
				Columns:    []*schema.Column{MetadataColumns[4]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// NodesColumns holds the columns for the "nodes" table.
	NodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"PACKAGE", "FILE"}},
		{Name: "name", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "file_name", Type: field.TypeString},
		{Name: "url_home", Type: field.TypeString},
		{Name: "url_download", Type: field.TypeString},
		{Name: "licenses", Type: field.TypeString},
		{Name: "license_concluded", Type: field.TypeString},
		{Name: "license_comments", Type: field.TypeString},
		{Name: "copyright", Type: field.TypeString},
		{Name: "source_info", Type: field.TypeString},
		{Name: "comment", Type: field.TypeString},
		{Name: "summary", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "attribution", Type: field.TypeString},
		{Name: "file_types", Type: field.TypeString},
		{Name: "primary_purpose", Type: field.TypeEnum, Enums: []string{"UNKNOWN_PURPOSE", "APPLICATION", "ARCHIVE", "BOM", "CONFIGURATION", "CONTAINER", "DATA", "DEVICE", "DEVICE_DRIVER", "DOCUMENTATION", "EVIDENCE", "EXECUTABLE", "FILE", "FIRMWARE", "FRAMEWORK", "INSTALL", "LIBRARY", "MACHINE_LEARNING_MODEL", "MANIFEST", "MODEL", "MODULE", "OPERATING_SYSTEM", "OTHER", "PATCH", "PLATFORM", "REQUIREMENT", "SOURCE", "SPECIFICATION", "TEST"}},
		{Name: "node_list_nodes", Type: field.TypeInt},
	}
	// NodesTable holds the schema information for the "nodes" table.
	NodesTable = &schema.Table{
		Name:       "nodes",
		Columns:    NodesColumns,
		PrimaryKey: []*schema.Column{NodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "nodes_node_lists_nodes",
				Columns:    []*schema.Column{NodesColumns[18]},
				RefColumns: []*schema.Column{NodeListsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// NodeListsColumns holds the columns for the "node_lists" table.
	NodeListsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "root_elements", Type: field.TypeString},
		{Name: "document_node_list", Type: field.TypeInt, Unique: true},
	}
	// NodeListsTable holds the schema information for the "node_lists" table.
	NodeListsTable = &schema.Table{
		Name:       "node_lists",
		Columns:    NodeListsColumns,
		PrimaryKey: []*schema.Column{NodeListsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_lists_documents_node_list",
				Columns:    []*schema.Column{NodeListsColumns[2]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PersonsColumns holds the columns for the "persons" table.
	PersonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_org", Type: field.TypeBool},
		{Name: "email", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "metadata_authors", Type: field.TypeString, Nullable: true},
		{Name: "node_suppliers", Type: field.TypeString, Nullable: true},
		{Name: "node_originators", Type: field.TypeString, Nullable: true},
	}
	// PersonsTable holds the schema information for the "persons" table.
	PersonsTable = &schema.Table{
		Name:       "persons",
		Columns:    PersonsColumns,
		PrimaryKey: []*schema.Column{PersonsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "persons_metadata_authors",
				Columns:    []*schema.Column{PersonsColumns[6]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "persons_nodes_suppliers",
				Columns:    []*schema.Column{PersonsColumns[7]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "persons_nodes_originators",
				Columns:    []*schema.Column{PersonsColumns[8]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TimestampsColumns holds the columns for the "timestamps" table.
	TimestampsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "metadata_date", Type: field.TypeString, Nullable: true},
		{Name: "node_release_date", Type: field.TypeString, Nullable: true},
		{Name: "node_build_date", Type: field.TypeString, Nullable: true},
		{Name: "node_valid_until_date", Type: field.TypeString, Nullable: true},
	}
	// TimestampsTable holds the schema information for the "timestamps" table.
	TimestampsTable = &schema.Table{
		Name:       "timestamps",
		Columns:    TimestampsColumns,
		PrimaryKey: []*schema.Column{TimestampsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "timestamps_metadata_date",
				Columns:    []*schema.Column{TimestampsColumns[1]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "timestamps_nodes_release_date",
				Columns:    []*schema.Column{TimestampsColumns[2]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "timestamps_nodes_build_date",
				Columns:    []*schema.Column{TimestampsColumns[3]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "timestamps_nodes_valid_until_date",
				Columns:    []*schema.Column{TimestampsColumns[4]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ToolsColumns holds the columns for the "tools" table.
	ToolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "vendor", Type: field.TypeString},
		{Name: "metadata_tools", Type: field.TypeString, Nullable: true},
	}
	// ToolsTable holds the schema information for the "tools" table.
	ToolsTable = &schema.Table{
		Name:       "tools",
		Columns:    ToolsColumns,
		PrimaryKey: []*schema.Column{ToolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tools_metadata_tools",
				Columns:    []*schema.Column{ToolsColumns[4]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ExternalReferenceHashesColumns holds the columns for the "external_reference_hashes" table.
	ExternalReferenceHashesColumns = []*schema.Column{
		{Name: "external_reference_id", Type: field.TypeInt},
		{Name: "hashes_entry_id", Type: field.TypeInt},
	}
	// ExternalReferenceHashesTable holds the schema information for the "external_reference_hashes" table.
	ExternalReferenceHashesTable = &schema.Table{
		Name:       "external_reference_hashes",
		Columns:    ExternalReferenceHashesColumns,
		PrimaryKey: []*schema.Column{ExternalReferenceHashesColumns[0], ExternalReferenceHashesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "external_reference_hashes_external_reference_id",
				Columns:    []*schema.Column{ExternalReferenceHashesColumns[0]},
				RefColumns: []*schema.Column{ExternalReferencesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "external_reference_hashes_hashes_entry_id",
				Columns:    []*schema.Column{ExternalReferenceHashesColumns[1]},
				RefColumns: []*schema.Column{HashesEntriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NodeIdentifiersColumns holds the columns for the "node_identifiers" table.
	NodeIdentifiersColumns = []*schema.Column{
		{Name: "node_id", Type: field.TypeString},
		{Name: "identifiers_entry_id", Type: field.TypeInt},
	}
	// NodeIdentifiersTable holds the schema information for the "node_identifiers" table.
	NodeIdentifiersTable = &schema.Table{
		Name:       "node_identifiers",
		Columns:    NodeIdentifiersColumns,
		PrimaryKey: []*schema.Column{NodeIdentifiersColumns[0], NodeIdentifiersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_identifiers_node_id",
				Columns:    []*schema.Column{NodeIdentifiersColumns[0]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "node_identifiers_identifiers_entry_id",
				Columns:    []*schema.Column{NodeIdentifiersColumns[1]},
				RefColumns: []*schema.Column{IdentifiersEntriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NodeHashesColumns holds the columns for the "node_hashes" table.
	NodeHashesColumns = []*schema.Column{
		{Name: "node_id", Type: field.TypeString},
		{Name: "hashes_entry_id", Type: field.TypeInt},
	}
	// NodeHashesTable holds the schema information for the "node_hashes" table.
	NodeHashesTable = &schema.Table{
		Name:       "node_hashes",
		Columns:    NodeHashesColumns,
		PrimaryKey: []*schema.Column{NodeHashesColumns[0], NodeHashesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_hashes_node_id",
				Columns:    []*schema.Column{NodeHashesColumns[0]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "node_hashes_hashes_entry_id",
				Columns:    []*schema.Column{NodeHashesColumns[1]},
				RefColumns: []*schema.Column{HashesEntriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PersonContactsColumns holds the columns for the "person_contacts" table.
	PersonContactsColumns = []*schema.Column{
		{Name: "person_id", Type: field.TypeInt},
		{Name: "contact_id", Type: field.TypeInt},
	}
	// PersonContactsTable holds the schema information for the "person_contacts" table.
	PersonContactsTable = &schema.Table{
		Name:       "person_contacts",
		Columns:    PersonContactsColumns,
		PrimaryKey: []*schema.Column{PersonContactsColumns[0], PersonContactsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "person_contacts_person_id",
				Columns:    []*schema.Column{PersonContactsColumns[0]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "person_contacts_contact_id",
				Columns:    []*schema.Column{PersonContactsColumns[1]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TimestampDateColumns holds the columns for the "timestamp_date" table.
	TimestampDateColumns = []*schema.Column{
		{Name: "timestamp_id", Type: field.TypeInt},
		{Name: "date_id", Type: field.TypeInt},
	}
	// TimestampDateTable holds the schema information for the "timestamp_date" table.
	TimestampDateTable = &schema.Table{
		Name:       "timestamp_date",
		Columns:    TimestampDateColumns,
		PrimaryKey: []*schema.Column{TimestampDateColumns[0], TimestampDateColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "timestamp_date_timestamp_id",
				Columns:    []*schema.Column{TimestampDateColumns[0]},
				RefColumns: []*schema.Column{TimestampsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "timestamp_date_date_id",
				Columns:    []*schema.Column{TimestampDateColumns[1]},
				RefColumns: []*schema.Column{TimestampsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DocumentsTable,
		DocumentTypesTable,
		EdgesTable,
		ExternalReferencesTable,
		HashesEntriesTable,
		IdentifiersEntriesTable,
		MetadataTable,
		NodesTable,
		NodeListsTable,
		PersonsTable,
		TimestampsTable,
		ToolsTable,
		ExternalReferenceHashesTable,
		NodeIdentifiersTable,
		NodeHashesTable,
		PersonContactsTable,
		TimestampDateTable,
	}
)

func init() {
	DocumentTypesTable.ForeignKeys[0].RefTable = MetadataTable
	EdgesTable.ForeignKeys[0].RefTable = NodeListsTable
	ExternalReferencesTable.ForeignKeys[0].RefTable = NodesTable
	MetadataTable.ForeignKeys[0].RefTable = DocumentsTable
	NodesTable.ForeignKeys[0].RefTable = NodeListsTable
	NodeListsTable.ForeignKeys[0].RefTable = DocumentsTable
	PersonsTable.ForeignKeys[0].RefTable = MetadataTable
	PersonsTable.ForeignKeys[1].RefTable = NodesTable
	PersonsTable.ForeignKeys[2].RefTable = NodesTable
	TimestampsTable.ForeignKeys[0].RefTable = MetadataTable
	TimestampsTable.ForeignKeys[1].RefTable = NodesTable
	TimestampsTable.ForeignKeys[2].RefTable = NodesTable
	TimestampsTable.ForeignKeys[3].RefTable = NodesTable
	ToolsTable.ForeignKeys[0].RefTable = MetadataTable
	ExternalReferenceHashesTable.ForeignKeys[0].RefTable = ExternalReferencesTable
	ExternalReferenceHashesTable.ForeignKeys[1].RefTable = HashesEntriesTable
	NodeIdentifiersTable.ForeignKeys[0].RefTable = NodesTable
	NodeIdentifiersTable.ForeignKeys[1].RefTable = IdentifiersEntriesTable
	NodeHashesTable.ForeignKeys[0].RefTable = NodesTable
	NodeHashesTable.ForeignKeys[1].RefTable = HashesEntriesTable
	PersonContactsTable.ForeignKeys[0].RefTable = PersonsTable
	PersonContactsTable.ForeignKeys[1].RefTable = PersonsTable
	TimestampDateTable.ForeignKeys[0].RefTable = TimestampsTable
	TimestampDateTable.ForeignKeys[1].RefTable = TimestampsTable
}
