# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import fog_pb2 as fog__pb2

GRPC_GENERATED_VERSION = '1.64.1'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.65.0'
SCHEDULED_RELEASE_DATE = 'June 25, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in fog_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class SensorServiceStub(object):
    """SensorService runs on the edge device and
    streams sensor data to the EdgeService.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.StreamData = channel.unary_stream(
                '/fogpb.SensorService/StreamData',
                request_serializer=fog__pb2.StreamDataRequest.SerializeToString,
                response_deserializer=fog__pb2.StreamDataResponse.FromString,
                _registered_method=True)


class SensorServiceServicer(object):
    """SensorService runs on the edge device and
    streams sensor data to the EdgeService.
    """

    def StreamData(self, request, context):
        """StreamData will be called by the EdgeService
        to initiate a stream of sensor data.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SensorServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'StreamData': grpc.unary_stream_rpc_method_handler(
                    servicer.StreamData,
                    request_deserializer=fog__pb2.StreamDataRequest.FromString,
                    response_serializer=fog__pb2.StreamDataResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'fogpb.SensorService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('fogpb.SensorService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class SensorService(object):
    """SensorService runs on the edge device and
    streams sensor data to the EdgeService.
    """

    @staticmethod
    def StreamData(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/fogpb.SensorService/StreamData',
            fog__pb2.StreamDataRequest.SerializeToString,
            fog__pb2.StreamDataResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)


class EdgeServiceStub(object):
    """EdgeService runs on the edge device and
    receives sensor data from the SensorService.
    Removed the ProcessDataStream RPC from EdgeService
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """


class EdgeServiceServicer(object):
    """EdgeService runs on the edge device and
    receives sensor data from the SensorService.
    Removed the ProcessDataStream RPC from EdgeService
    """


def add_EdgeServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'fogpb.EdgeService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('fogpb.EdgeService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class EdgeService(object):
    """EdgeService runs on the edge device and
    receives sensor data from the SensorService.
    Removed the ProcessDataStream RPC from EdgeService
    """


class CloudServiceStub(object):
    """CloudService runs on the cloud and processes
    aggregated sensor data.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ProcessDataStream = channel.stream_stream(
                '/fogpb.CloudService/ProcessDataStream',
                request_serializer=fog__pb2.ProcessDataRequest.SerializeToString,
                response_deserializer=fog__pb2.UpdatePositionResponse.FromString,
                _registered_method=True)


class CloudServiceServicer(object):
    """CloudService runs on the cloud and processes
    aggregated sensor data.
    """

    def ProcessDataStream(self, request_iterator, context):
        """Bidirectional streaming endpoint for processing data
        and sending back position updates to the EdgeService.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CloudServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ProcessDataStream': grpc.stream_stream_rpc_method_handler(
                    servicer.ProcessDataStream,
                    request_deserializer=fog__pb2.ProcessDataRequest.FromString,
                    response_serializer=fog__pb2.UpdatePositionResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'fogpb.CloudService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('fogpb.CloudService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class CloudService(object):
    """CloudService runs on the cloud and processes
    aggregated sensor data.
    """

    @staticmethod
    def ProcessDataStream(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(
            request_iterator,
            target,
            '/fogpb.CloudService/ProcessDataStream',
            fog__pb2.ProcessDataRequest.SerializeToString,
            fog__pb2.UpdatePositionResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
