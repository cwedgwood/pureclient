# DirectoryPerformance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**BytesPerOp** | **int64** | The average I/O size for both read and write (all) operations. | [optional] [default to null]
**BytesPerRead** | **int64** | The average I/O size per read, measured in bytes. | [optional] [default to null]
**BytesPerWrite** | **int64** | The average I/O size per write, measured in bytes. | [optional] [default to null]
**OthersPerSec** | **int64** | The number of other requests processed per second. | [optional] [default to null]
**ReadBytesPerSec** | **int64** | The number of bytes read per second. | [optional] [default to null]
**ReadsPerSec** | **int64** | The number of read requests processed per second. | [optional] [default to null]
**Time** | **int64** | The time when the sample performance data was taken. Measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**UsecPerOtherOp** | **int64** | The average time it takes the array to process an I/O other request, measured in microseconds. The average time does not include SAN time, queue time, or QoS rate limit time. | [optional] [default to null]
**UsecPerReadOp** | **int64** | The average time it takes the array to process an I/O read request, measured in microseconds. The average time does not include SAN time, queue time, or QoS rate limit time. | [optional] [default to null]
**UsecPerWriteOp** | **int64** | The average time it takes the array to process an I/O write request, measured in microseconds. The average time does not include SAN time, queue time, or QoS rate limit time. | [optional] [default to null]
**WriteBytesPerSec** | **int64** | The number of bytes written per second. | [optional] [default to null]
**WritesPerSec** | **int64** | The number of write requests processed per second. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

