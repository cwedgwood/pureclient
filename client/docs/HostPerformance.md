# HostPerformance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesPerMirroredWrite** | **int64** | The average I/O size per mirrored write. Measured in bytes. | [optional] [default to null]
**BytesPerOp** | **int64** | The average I/O size for both read and write (all) operations. | [optional] [default to null]
**BytesPerRead** | **int64** | The average I/O size per read. Measured in bytes. | [optional] [default to null]
**BytesPerWrite** | **int64** | The average I/O size per write. Measured in bytes. | [optional] [default to null]
**MirroredWriteBytesPerSec** | **int64** | The number of mirrored bytes written per second. | [optional] [default to null]
**MirroredWritesPerSec** | **int64** | The number of mirrored writes per second. | [optional] [default to null]
**QosRateLimitUsecPerMirroredWriteOp** | **int64** | The average time it takes the array to process a mirrored I/O write request. Measured in microseconds. | [optional] [default to null]
**QosRateLimitUsecPerReadOp** | **int64** | The average time spent waiting due to QoS rate limiting for a read request. Measured in microseconds. | [optional] [default to null]
**QosRateLimitUsecPerWriteOp** | **int64** | The average time that a write I/O request spends waiting as a result of the volume reaching its QoS bandwidth limit. Measured in microseconds. | [optional] [default to null]
**QueueUsecPerMirroredWriteOp** | **int64** | The average time that a mirrored write I/O request spends in the array waiting to be served. Measured in microseconds. | [optional] [default to null]
**QueueUsecPerReadOp** | **int64** | The average time that a read I/O request spends in the array waiting to be served. Measured in microseconds. | [optional] [default to null]
**QueueUsecPerWriteOp** | **int64** | The average time that a write I/O request spends in the array waiting to be served. Measured in microseconds. | [optional] [default to null]
**ReadBytesPerSec** | **int64** | The number of bytes read per second. | [optional] [default to null]
**ReadsPerSec** | **int64** | The number of read requests processed per second. | [optional] [default to null]
**SanUsecPerMirroredWriteOp** | **int64** | The average time required to transfer data from the initiator to the array for a mirrored write request. Measured in microseconds. | [optional] [default to null]
**SanUsecPerReadOp** | **int64** | The average time required to transfer data from the array to the initiator for a read request. Measured in microseconds. | [optional] [default to null]
**SanUsecPerWriteOp** | **int64** | The average time required to transfer data from the initiator to the array for a write request. Measured in microseconds. | [optional] [default to null]
**ServiceUsecPerMirroredWriteOp** | **int64** | The average time required for the array to service a mirrored write request. Measured in microseconds. | [optional] [default to null]
**ServiceUsecPerReadOp** | **int64** | The average time required for the array to service a read request. Measured in microseconds. | [optional] [default to null]
**ServiceUsecPerWriteOp** | **int64** | The average time required for the array to service a write request. Measured in microseconds. | [optional] [default to null]
**Time** | **int64** | The time when the sample performance data was taken. Measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**UsecPerMirroredWriteOp** | **int64** | The average time it takes the array to process a mirrored I/O write request. Measured in microseconds. The average time does not include SAN time, queue time, or QoS rate limit time. | [optional] [default to null]
**UsecPerReadOp** | **int64** | The average time it takes the array to process an I/O read request. Measured in microseconds. The average time does not include SAN time, queue time, or QoS rate limit time. | [optional] [default to null]
**UsecPerWriteOp** | **int64** | The average time it takes the array to process an I/O write request. Measured in microseconds. The average time does not include SAN time, queue time, or QoS rate limit time. | [optional] [default to null]
**WriteBytesPerSec** | **int64** | The number of bytes written per second. | [optional] [default to null]
**WritesPerSec** | **int64** | The number of write requests processed per second. | [optional] [default to null]
**ServiceUsecPerReadOpCacheReduction** | **float32** | The percentage reduction in &#x60;service_usec_per_read_op&#x60; due to data cache hits. For example, a value of 0.25 indicates that the value of &#x60;service_usec_per_read_op&#x60; is 25&amp;#37; lower than it would have been without any data cache hits. | [optional] [default to null]
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

