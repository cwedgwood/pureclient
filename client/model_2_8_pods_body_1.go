/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Model28PodsBody1 struct {
	// A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource.
	Id string `json:"id,omitempty"`
	// A user-specified name. The name must be locally unique and can be changed.
	Name string `json:"name,omitempty"`
	// If set to `true`, the pod has been destroyed and is pending eradication. The `time_remaining` value displays the amount of time left until the destroyed pod is permanently eradicated. A pod can only be destroyed if it is empty, so before destroying a pod, ensure all volumes and protection groups inside the pod have been either moved out of the pod or destroyed. A stretched pod cannot be destroyed unless you unstretch it first. Before the `time_remaining` period has elapsed, the destroyed pod can be recovered by setting `destroyed=false`. Once the `time_remaining` period has elapsed, the pod is permanently eradicated and can no longer be recovered.
	Destroyed bool `json:"destroyed,omitempty"`
	// Determines which array within a stretched pod should be given priority to stay online should the arrays ever lose contact with each other. The current array and any peer arrays that are connected to the current array for synchronous replication can be added to a pod for failover preference. By default, `failover_preferences=null`, meaning no arrays have been configured for failover preference. Enter multiple arrays in comma-separated format.
	FailoverPreferences []interface{} `json:"failover_preferences,omitempty"`
	// Sets the URL of the mediator for this pod, replacing the URL of the current mediator. By default, the Pure1 Cloud Mediator (`purestorage`) serves as the mediator.
	Mediator string `json:"mediator,omitempty"`
	// Patch `requested_promotion_state` to `demoted` to demote the pod so that it can be used as a link target for continuous replication between pods. Demoted pods do not accept write requests, and a destroyed version of the pod with `undo-demote` appended to the pod name is created on the array with the state of the pod when it was in the promoted state. Patch `requested_promotion_state` to `promoted` to start the process of promoting the pod. The `promotion_status` indicates when the pod has been successfully promoted. Promoted pods stop incorporating replicated data from the source pod and start accepting write requests. The replication process does not stop when the source pod continues replicating data to the pod. The space consumed by the unique replicated data is tracked by the `space.journal` field of the pod.
	RequestedPromotionState string `json:"requested_promotion_state,omitempty"`
}
