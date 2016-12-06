package hypercloud

type ConsoleSession struct {
	client *ApiClient
}
type Disk struct {
	client *ApiClient
}
type Instance struct {
	client *ApiClient
}
type IpAddress struct {
	client *ApiClient
}
type Network struct {
	client *ApiClient
}
type PerformanceTier struct {
	client *ApiClient
}
type PublicKey struct {
	client *ApiClient
}
type Region struct {
	client *ApiClient
}
type Template struct {
	client *ApiClient
}

func (resource *ConsoleSession) Show(console_session_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/console_sessions/" + console_session_id + ""
	return resource.client.MapRequest(url, "GET", nil)
}

func (resource *Disk) Show(disk_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/disks/" + disk_id + ""
	return resource.client.MapRequest(url, "GET", nil)
}
func (resource *Disk) Create(body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/disks"
	return resource.client.MapRequest(url, "POST", body)
}
func (resource *Disk) Delete(disk_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/disks/" + disk_id + ""
	return resource.client.MapRequest(url, "DELETE", nil)
}
func (resource *Disk) State(disk_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/disks/" + disk_id + "/state"
	return resource.client.MapRequest(url, "GET", nil)
}
func (resource *Disk) List() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/disks"
	return resource.client.ArrayRequest(url, "GET", nil)
}
func (resource *Disk) Update(disk_id string, body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/disks/" + disk_id + ""
	return resource.client.MapRequest(url, "PUT", body)
}
func (resource *Disk) Resize(disk_id string, body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/disks/" + disk_id + "/resize"
	return resource.client.MapRequest(url, "POST", body)
}
func (resource *Disk) Clone(disk_id string, body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/disks/" + disk_id + "/clone"
	return resource.client.MapRequest(url, "POST", body)
}

func (resource *Instance) Create_basic(body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances"
	return resource.client.MapRequest(url, "POST", body)
}
func (resource *Instance) Create_advanced(body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/assemble"
	return resource.client.MapRequest(url, "POST", body)
}
func (resource *Instance) Delete(instance_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + ""
	return resource.client.MapRequest(url, "DELETE", nil)
}
func (resource *Instance) Show(instance_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + ""
	return resource.client.MapRequest(url, "GET", nil)
}
func (resource *Instance) List() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances"
	return resource.client.ArrayRequest(url, "GET", nil)
}
func (resource *Instance) Update(instance_id string, body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + ""
	return resource.client.MapRequest(url, "PUT", body)
}
func (resource *Instance) State(instance_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/state"
	return resource.client.MapRequest(url, "GET", nil)
}
func (resource *Instance) Note(instance_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/note"
	return resource.client.MapRequest(url, "GET", nil)
}
func (resource *Instance) Start(instance_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/start"
	return resource.client.MapRequest(url, "POST", nil)
}
func (resource *Instance) Stop(instance_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/stop"
	return resource.client.MapRequest(url, "POST", nil)
}
func (resource *Instance) Remote_access(instance_id string, body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/remote_access"
	return resource.client.MapRequest(url, "POST", body)
}
func (resource *Instance) Update_disks(instance_id string, body interface{}) (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/disks"
	return resource.client.ArrayRequest(url, "PUT", body)
}
func (resource *Instance) Update_public_keys(instance_id string, body interface{}) (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/public_keys"
	return resource.client.ArrayRequest(url, "PUT", body)
}
func (resource *Instance) Update_networking(instance_id string, body interface{}) (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/network_adapters"
	return resource.client.ArrayRequest(url, "PUT", body)
}
func (resource *Instance) Update_high_availability(instance_id string, body interface{}) (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/availability_group"
	return resource.client.ArrayRequest(url, "PUT", body)
}
func (resource *Instance) Get_context(instance_id string) (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/context"
	return resource.client.ArrayRequest(url, "GET", nil)
}
func (resource *Instance) Set_context(instance_id string, body interface{}) (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/context"
	return resource.client.ArrayRequest(url, "POST", body)
}
func (resource *Instance) Update_context(instance_id string, body interface{}) (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/context"
	return resource.client.ArrayRequest(url, "PUT", body)
}
func (resource *Instance) Delete_context_variable(instance_id string, instance_context_key string) (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/instances/" + instance_id + "/context/" + instance_context_key + ""
	return resource.client.ArrayRequest(url, "DELETE", nil)
}

func (resource *IpAddress) Allocate(body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/ip_addresses"
	return resource.client.MapRequest(url, "POST", body)
}
func (resource *IpAddress) Deallocate(ip_address_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/ip_addresses/" + ip_address_id + ""
	return resource.client.MapRequest(url, "DELETE", nil)
}
func (resource *IpAddress) List() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/ip_addresses"
	return resource.client.ArrayRequest(url, "GET", nil)
}
func (resource *IpAddress) List_public() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/ip_addresses/public"
	return resource.client.ArrayRequest(url, "GET", nil)
}
func (resource *IpAddress) List_private() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/ip_addresses/private"
	return resource.client.ArrayRequest(url, "GET", nil)
}
func (resource *IpAddress) Show(ip_address_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/ip_addresses/" + ip_address_id + ""
	return resource.client.MapRequest(url, "GET", nil)
}
func (resource *IpAddress) Update(ip_address_id string, body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/ip_addresses/" + ip_address_id + ""
	return resource.client.MapRequest(url, "PUT", body)
}

func (resource *Network) Create(body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/networks"
	return resource.client.MapRequest(url, "POST", body)
}
func (resource *Network) Delete(network_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/networks/" + network_id + ""
	return resource.client.MapRequest(url, "DELETE", nil)
}
func (resource *Network) List() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/networks"
	return resource.client.ArrayRequest(url, "GET", nil)
}
func (resource *Network) List_public() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/networks/public"
	return resource.client.ArrayRequest(url, "GET", nil)
}
func (resource *Network) List_private() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/networks/private"
	return resource.client.ArrayRequest(url, "GET", nil)
}
func (resource *Network) Show(network_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/networks/" + network_id + ""
	return resource.client.MapRequest(url, "GET", nil)
}
func (resource *Network) Update(network_id string, body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/networks/" + network_id + ""
	return resource.client.MapRequest(url, "PUT", body)
}

func (resource *PerformanceTier) List_instance() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/performance_tiers/instances"
	return resource.client.ArrayRequest(url, "GET", nil)
}
func (resource *PerformanceTier) List_disk() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/performance_tiers/disks"
	return resource.client.ArrayRequest(url, "GET", nil)
}

func (resource *PublicKey) Create(body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/public_keys"
	return resource.client.MapRequest(url, "POST", body)
}
func (resource *PublicKey) Delete(public_key_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/public_keys/" + public_key_id + ""
	return resource.client.MapRequest(url, "DELETE", nil)
}
func (resource *PublicKey) Show(public_key_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/public_keys/" + public_key_id + ""
	return resource.client.MapRequest(url, "GET", nil)
}
func (resource *PublicKey) List() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/public_keys"
	return resource.client.ArrayRequest(url, "GET", nil)
}
func (resource *PublicKey) Update(public_key_id string, body interface{}) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/public_keys/" + public_key_id + ""
	return resource.client.MapRequest(url, "PUT", body)
}

func (resource *Region) Show(region_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/regions/" + region_id + ""
	return resource.client.MapRequest(url, "GET", nil)
}
func (resource *Region) List() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/regions"
	return resource.client.ArrayRequest(url, "GET", nil)
}

func (resource *Template) Show(template_id string) (status int, rbody map[string]interface{}, err error) {
	url := resource.client.apiPath + "/templates/" + template_id + ""
	return resource.client.MapRequest(url, "GET", nil)
}
func (resource *Template) List() (status int, rbody []map[string]interface{}, err error) {
	url := resource.client.apiPath + "/templates"
	return resource.client.ArrayRequest(url, "GET", nil)
}
