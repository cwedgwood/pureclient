/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

// Ethernet network interface statistics.
type NetworkInterfacePerformanceEth struct {
	// The sum of unspecified reception and transmission errors per second.
	OtherErrorsPerSec int64 `json:"other_errors_per_sec,omitempty"`
	// Bytes received per second.
	ReceivedBytesPerSec int64 `json:"received_bytes_per_sec,omitempty"`
	// Reception CRC errors per second.
	ReceivedCrcErrorsPerSec int64 `json:"received_crc_errors_per_sec,omitempty"`
	// Received packet frame errors per second.
	ReceivedFrameErrorsPerSec int64 `json:"received_frame_errors_per_sec,omitempty"`
	// Packets received per second.
	ReceivedPacketsPerSec int64 `json:"received_packets_per_sec,omitempty"`
	// The sum of all reception and transmission errors per second.
	TotalErrorsPerSec int64 `json:"total_errors_per_sec,omitempty"`
	// Bytes transmitted per second.
	TransmittedBytesPerSec int64 `json:"transmitted_bytes_per_sec,omitempty"`
	// Transmission carrier errors per second.
	TransmittedCarrierErrorsPerSec int64 `json:"transmitted_carrier_errors_per_sec,omitempty"`
	// Transmitted packets dropped per second.
	TransmittedDroppedErrorsPerSec int64 `json:"transmitted_dropped_errors_per_sec,omitempty"`
	// Packets transmitted per second.
	TransmittedPacketsPerSec int64 `json:"transmitted_packets_per_sec,omitempty"`
}
