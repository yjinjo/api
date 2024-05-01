# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.inventory.v1 import cloud_service_type_pb2 as spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2

GRPC_GENERATED_VERSION = '1.63.0'
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
        + f' but the generated code in spaceone/api/inventory/v1/cloud_service_type_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class CloudServiceTypeStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceType/create',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CreateCloudServiceTypeRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeInfo.FromString,
                _registered_method=True)
        self.update = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceType/update',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.UpdateCloudServiceTypeRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeInfo.FromString,
                _registered_method=True)
        self.delete = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceType/delete',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.get = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceType/get',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeInfo.FromString,
                _registered_method=True)
        self.list = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceType/list',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypesInfo.FromString,
                _registered_method=True)
        self.stat = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceType/stat',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                _registered_method=True)


class CloudServiceTypeServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Creates a new CloudServiceType. You must specify the `name`, `provider`, and `group` parameters to create a CloudServiceType. One or several CloudServiceTypes exist in a specific `group`, and each CloudServiceType is identified by the `name` parameter.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Updates a specific CloudServiceType. You can make changes in CloudServiceType settings, except for `name`, `provider` and `group`. In particular, you can set the CloudServiceType's priority in a `group`.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Deletes a specific CloudServiceType. You must specify the `cloud_service_type_id` of the CloudServiceType to delete.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Gets a specific CloudServiceType. Prints detailed information about the CloudServiceType.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Gets a list of all CloudServiceTypes. You can use a query to get a filtered list of CloudServiceTypes.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CloudServiceTypeServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CreateCloudServiceTypeRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.UpdateCloudServiceTypeRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypesInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.inventory.v1.CloudServiceType', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CloudServiceType(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.inventory.v1.CloudServiceType/create',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CreateCloudServiceTypeRequest.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.inventory.v1.CloudServiceType/update',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.UpdateCloudServiceTypeRequest.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.inventory.v1.CloudServiceType/delete',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.inventory.v1.CloudServiceType/get',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeRequest.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def list(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.inventory.v1.CloudServiceType/list',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeQuery.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypesInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def stat(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.inventory.v1.CloudServiceType/stat',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__type__pb2.CloudServiceTypeStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
