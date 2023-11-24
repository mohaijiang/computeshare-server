// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AgentsColumns holds the columns for the "agents" table.
	AgentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "peer_id", Type: field.TypeString, Unique: true},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "last_update_time", Type: field.TypeTime},
	}
	// AgentsTable holds the schema information for the "agents" table.
	AgentsTable = &schema.Table{
		Name:       "agents",
		Columns:    AgentsColumns,
		PrimaryKey: []*schema.Column{AgentsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "agent_id",
				Unique:  true,
				Columns: []*schema.Column{AgentsColumns[0]},
			},
		},
	}
	// ComputeImagesColumns holds the columns for the "compute_images" table.
	ComputeImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "image", Type: field.TypeString},
		{Name: "tag", Type: field.TypeString},
		{Name: "port", Type: field.TypeInt32},
		{Name: "command", Type: field.TypeString},
	}
	// ComputeImagesTable holds the schema information for the "compute_images" table.
	ComputeImagesTable = &schema.Table{
		Name:       "compute_images",
		Columns:    ComputeImagesColumns,
		PrimaryKey: []*schema.Column{ComputeImagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "computeimage_id",
				Unique:  true,
				Columns: []*schema.Column{ComputeImagesColumns[0]},
			},
		},
	}
	// ComputeInstancesColumns holds the columns for the "compute_instances" table.
	ComputeInstancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "owner", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "core", Type: field.TypeString},
		{Name: "memory", Type: field.TypeString},
		{Name: "image", Type: field.TypeString},
		{Name: "port", Type: field.TypeString, Nullable: true},
		{Name: "expiration_time", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8},
		{Name: "container_id", Type: field.TypeString, Nullable: true},
		{Name: "agent_id", Type: field.TypeString, Nullable: true},
		{Name: "command", Type: field.TypeString, Nullable: true},
	}
	// ComputeInstancesTable holds the schema information for the "compute_instances" table.
	ComputeInstancesTable = &schema.Table{
		Name:       "compute_instances",
		Columns:    ComputeInstancesColumns,
		PrimaryKey: []*schema.Column{ComputeInstancesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "computeinstance_id",
				Unique:  true,
				Columns: []*schema.Column{ComputeInstancesColumns[0]},
			},
			{
				Name:    "computeinstance_owner",
				Unique:  false,
				Columns: []*schema.Column{ComputeInstancesColumns[1]},
			},
		},
	}
	// ComputeSpecsColumns holds the columns for the "compute_specs" table.
	ComputeSpecsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "core", Type: field.TypeString},
		{Name: "memory", Type: field.TypeString},
	}
	// ComputeSpecsTable holds the schema information for the "compute_specs" table.
	ComputeSpecsTable = &schema.Table{
		Name:       "compute_specs",
		Columns:    ComputeSpecsColumns,
		PrimaryKey: []*schema.Column{ComputeSpecsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "computespec_id",
				Unique:  true,
				Columns: []*schema.Column{ComputeSpecsColumns[0]},
			},
		},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt32, Nullable: true},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:       "employees",
		Columns:    EmployeesColumns,
		PrimaryKey: []*schema.Column{EmployeesColumns[0]},
	}
	// GatewaysColumns holds the columns for the "gateways" table.
	GatewaysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "ip", Type: field.TypeString},
		{Name: "port", Type: field.TypeInt},
	}
	// GatewaysTable holds the schema information for the "gateways" table.
	GatewaysTable = &schema.Table{
		Name:       "gateways",
		Columns:    GatewaysColumns,
		PrimaryKey: []*schema.Column{GatewaysColumns[0]},
	}
	// GatewayPortsColumns holds the columns for the "gateway_ports" table.
	GatewayPortsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "fk_gateway_id", Type: field.TypeUUID},
		{Name: "port", Type: field.TypeInt64},
		{Name: "is_use", Type: field.TypeBool, Default: false},
	}
	// GatewayPortsTable holds the schema information for the "gateway_ports" table.
	GatewayPortsTable = &schema.Table{
		Name:       "gateway_ports",
		Columns:    GatewayPortsColumns,
		PrimaryKey: []*schema.Column{GatewayPortsColumns[0]},
	}
	// NetworkMappingsColumns holds the columns for the "network_mappings" table.
	NetworkMappingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "fk_gateway_id", Type: field.TypeUUID},
		{Name: "fk_computer_id", Type: field.TypeUUID},
		{Name: "gateway_port", Type: field.TypeInt},
		{Name: "computer_port", Type: field.TypeInt},
		{Name: "status", Type: field.TypeInt, Default: 0},
	}
	// NetworkMappingsTable holds the schema information for the "network_mappings" table.
	NetworkMappingsTable = &schema.Table{
		Name:       "network_mappings",
		Columns:    NetworkMappingsColumns,
		PrimaryKey: []*schema.Column{NetworkMappingsColumns[0]},
	}
	// ScriptsColumns holds the columns for the "scripts" table.
	ScriptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "task_number", Type: field.TypeInt32},
		{Name: "script_name", Type: field.TypeString},
		{Name: "file_address", Type: field.TypeString},
		{Name: "script_content", Type: field.TypeString, Size: 2147483647},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// ScriptsTable holds the schema information for the "scripts" table.
	ScriptsTable = &schema.Table{
		Name:       "scripts",
		Columns:    ScriptsColumns,
		PrimaryKey: []*schema.Column{ScriptsColumns[0]},
	}
	// ScriptExecutionRecordsColumns holds the columns for the "script_execution_records" table.
	ScriptExecutionRecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "fk_script_id", Type: field.TypeInt32},
		{Name: "script_content", Type: field.TypeString, Size: 2147483647},
		{Name: "task_number", Type: field.TypeInt32},
		{Name: "script_name", Type: field.TypeString},
		{Name: "file_address", Type: field.TypeString},
		{Name: "execute_state", Type: field.TypeInt32},
		{Name: "execute_result", Type: field.TypeString, Size: 2147483647},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "script_script_execution_records", Type: field.TypeInt32, Nullable: true},
	}
	// ScriptExecutionRecordsTable holds the schema information for the "script_execution_records" table.
	ScriptExecutionRecordsTable = &schema.Table{
		Name:       "script_execution_records",
		Columns:    ScriptExecutionRecordsColumns,
		PrimaryKey: []*schema.Column{ScriptExecutionRecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "script_execution_records_scripts_scriptExecutionRecords",
				Columns:    []*schema.Column{ScriptExecutionRecordsColumns[11]},
				RefColumns: []*schema.Column{ScriptsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StoragesColumns holds the columns for the "storages" table.
	StoragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "owner", Type: field.TypeString, Size: 50},
		{Name: "type", Type: field.TypeInt32, Default: 0},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "cid", Type: field.TypeString, Size: 80},
		{Name: "size", Type: field.TypeInt32},
		{Name: "last_modify", Type: field.TypeTime},
		{Name: "parent_id", Type: field.TypeString, Size: 80},
	}
	// StoragesTable holds the schema information for the "storages" table.
	StoragesTable = &schema.Table{
		Name:       "storages",
		Columns:    StoragesColumns,
		PrimaryKey: []*schema.Column{StoragesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "storage_owner",
				Unique:  false,
				Columns: []*schema.Column{StoragesColumns[1]},
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "agent_id", Type: field.TypeString, Size: 50},
		{Name: "cmd", Type: field.TypeInt32, Default: 0},
		{Name: "params", Type: field.TypeString, Size: 1024},
		{Name: "status", Type: field.TypeInt},
		{Name: "create_time", Type: field.TypeTime},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "task_agent_id",
				Unique:  false,
				Columns: []*schema.Column{TasksColumns[1]},
			},
			{
				Name:    "task_create_time",
				Unique:  false,
				Columns: []*schema.Column{TasksColumns[5]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "country_call_coding", Type: field.TypeString, Size: 8},
		{Name: "telephone_number", Type: field.TypeString, Size: 50},
		{Name: "password", Type: field.TypeString},
		{Name: "create_date", Type: field.TypeTime},
		{Name: "last_login_date", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString},
		{Name: "pwd_config", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[0]},
			},
			{
				Name:    "user_country_call_coding_telephone_number",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1], UsersColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AgentsTable,
		ComputeImagesTable,
		ComputeInstancesTable,
		ComputeSpecsTable,
		EmployeesTable,
		GatewaysTable,
		GatewayPortsTable,
		NetworkMappingsTable,
		ScriptsTable,
		ScriptExecutionRecordsTable,
		StoragesTable,
		TasksTable,
		UsersTable,
	}
)

func init() {
	ScriptExecutionRecordsTable.ForeignKeys[0].RefTable = ScriptsTable
}
