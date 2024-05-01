# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.file_manager.v1 import file_pb2 as spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2

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
        + f' but the generated code in spaceone/api/file_manager/v1/file_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class FileStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.add = channel.unary_unary(
                '/spaceone.api.file_manager.v1.File/add',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.CreateFileRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.FromString,
                _registered_method=True)
        self.update = channel.unary_unary(
                '/spaceone.api.file_manager.v1.File/update',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.UpdateFileRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.FromString,
                _registered_method=True)
        self.delete = channel.unary_unary(
                '/spaceone.api.file_manager.v1.File/delete',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.get_download_url = channel.unary_unary(
                '/spaceone.api.file_manager.v1.File/get_download_url',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.FromString,
                _registered_method=True)
        self.get = channel.unary_unary(
                '/spaceone.api.file_manager.v1.File/get',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.FromString,
                _registered_method=True)
        self.list = channel.unary_unary(
                '/spaceone.api.file_manager.v1.File/list',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileSearchQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FilesInfo.FromString,
                _registered_method=True)
        self.stat = channel.unary_unary(
                '/spaceone.api.file_manager.v1.File/stat',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                _registered_method=True)


class FileServicer(object):
    """Missing associated documentation comment in .proto file."""

    def add(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get_download_url(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_FileServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'add': grpc.unary_unary_rpc_method_handler(
                    servicer.add,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.CreateFileRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.UpdateFileRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get_download_url': grpc.unary_unary_rpc_method_handler(
                    servicer.get_download_url,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileSearchQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FilesInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.file_manager.v1.File', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class File(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def add(request,
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
            '/spaceone.api.file_manager.v1.File/add',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.CreateFileRequest.SerializeToString,
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.FromString,
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
            '/spaceone.api.file_manager.v1.File/update',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.UpdateFileRequest.SerializeToString,
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.FromString,
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
            '/spaceone.api.file_manager.v1.File/delete',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileRequest.SerializeToString,
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
    def get_download_url(request,
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
            '/spaceone.api.file_manager.v1.File/get_download_url',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileRequest.SerializeToString,
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.FromString,
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
            '/spaceone.api.file_manager.v1.File/get',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileRequest.SerializeToString,
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileInfo.FromString,
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
            '/spaceone.api.file_manager.v1.File/list',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileSearchQuery.SerializeToString,
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FilesInfo.FromString,
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
            '/spaceone.api.file_manager.v1.File/stat',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2.FileStatQuery.SerializeToString,
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
