# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
import groups_pb2 as groups__pb2
import robots_pb2 as robots__pb2


class robotsStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetAll = channel.unary_unary(
                '/robots.robots/GetAll',
                request_serializer=robots__pb2.Robots.SerializeToString,
                response_deserializer=robots__pb2.Robots.FromString,
                )
        self.GetGraph = channel.unary_unary(
                '/robots.robots/GetGraph',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=groups__pb2.Groups.FromString,
                )
        self.Get = channel.unary_unary(
                '/robots.robots/Get',
                request_serializer=robots__pb2.Robot.SerializeToString,
                response_deserializer=robots__pb2.Robot.FromString,
                )
        self.GetDirectusToken = channel.unary_unary(
                '/robots.robots/GetDirectusToken',
                request_serializer=robots__pb2.Robot.SerializeToString,
                response_deserializer=robots__pb2.DirectusToken.FromString,
                )
        self.GetMyDirectusToken = channel.unary_unary(
                '/robots.robots/GetMyDirectusToken',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=robots__pb2.DirectusToken.FromString,
                )
        self.GetByGroup = channel.unary_unary(
                '/robots.robots/GetByGroup',
                request_serializer=groups__pb2.Group.SerializeToString,
                response_deserializer=robots__pb2.Robots.FromString,
                )
        self.GetSAFile = channel.unary_unary(
                '/robots.robots/GetSAFile',
                request_serializer=robots__pb2.Robot.SerializeToString,
                response_deserializer=robots__pb2.SaFile.FromString,
                )
        self.Add = channel.unary_unary(
                '/robots.robots/Add',
                request_serializer=robots__pb2.Robot.SerializeToString,
                response_deserializer=robots__pb2.Robot.FromString,
                )
        self.SendToRobot = channel.unary_unary(
                '/robots.robots/SendToRobot',
                request_serializer=robots__pb2.RobotMessage.SerializeToString,
                response_deserializer=robots__pb2.RobotMessage.FromString,
                )
        self.SendToRobotWithEmail = channel.unary_unary(
                '/robots.robots/SendToRobotWithEmail',
                request_serializer=robots__pb2.RobotMessage.SerializeToString,
                response_deserializer=robots__pb2.RobotMessage.FromString,
                )
        self.Update = channel.unary_unary(
                '/robots.robots/Update',
                request_serializer=robots__pb2.Robot.SerializeToString,
                response_deserializer=robots__pb2.Robot.FromString,
                )
        self.UpdateMacAddress = channel.unary_unary(
                '/robots.robots/UpdateMacAddress',
                request_serializer=robots__pb2.Robot.SerializeToString,
                response_deserializer=robots__pb2.Robot.FromString,
                )


class robotsServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetAll(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetGraph(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetDirectusToken(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMyDirectusToken(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetByGroup(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetSAFile(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Add(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendToRobot(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendToRobotWithEmail(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateMacAddress(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_robotsServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetAll': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAll,
                    request_deserializer=robots__pb2.Robots.FromString,
                    response_serializer=robots__pb2.Robots.SerializeToString,
            ),
            'GetGraph': grpc.unary_unary_rpc_method_handler(
                    servicer.GetGraph,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=groups__pb2.Groups.SerializeToString,
            ),
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=robots__pb2.Robot.FromString,
                    response_serializer=robots__pb2.Robot.SerializeToString,
            ),
            'GetDirectusToken': grpc.unary_unary_rpc_method_handler(
                    servicer.GetDirectusToken,
                    request_deserializer=robots__pb2.Robot.FromString,
                    response_serializer=robots__pb2.DirectusToken.SerializeToString,
            ),
            'GetMyDirectusToken': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMyDirectusToken,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=robots__pb2.DirectusToken.SerializeToString,
            ),
            'GetByGroup': grpc.unary_unary_rpc_method_handler(
                    servicer.GetByGroup,
                    request_deserializer=groups__pb2.Group.FromString,
                    response_serializer=robots__pb2.Robots.SerializeToString,
            ),
            'GetSAFile': grpc.unary_unary_rpc_method_handler(
                    servicer.GetSAFile,
                    request_deserializer=robots__pb2.Robot.FromString,
                    response_serializer=robots__pb2.SaFile.SerializeToString,
            ),
            'Add': grpc.unary_unary_rpc_method_handler(
                    servicer.Add,
                    request_deserializer=robots__pb2.Robot.FromString,
                    response_serializer=robots__pb2.Robot.SerializeToString,
            ),
            'SendToRobot': grpc.unary_unary_rpc_method_handler(
                    servicer.SendToRobot,
                    request_deserializer=robots__pb2.RobotMessage.FromString,
                    response_serializer=robots__pb2.RobotMessage.SerializeToString,
            ),
            'SendToRobotWithEmail': grpc.unary_unary_rpc_method_handler(
                    servicer.SendToRobotWithEmail,
                    request_deserializer=robots__pb2.RobotMessage.FromString,
                    response_serializer=robots__pb2.RobotMessage.SerializeToString,
            ),
            'Update': grpc.unary_unary_rpc_method_handler(
                    servicer.Update,
                    request_deserializer=robots__pb2.Robot.FromString,
                    response_serializer=robots__pb2.Robot.SerializeToString,
            ),
            'UpdateMacAddress': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateMacAddress,
                    request_deserializer=robots__pb2.Robot.FromString,
                    response_serializer=robots__pb2.Robot.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'robots.robots', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class robots(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetAll(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/GetAll',
            robots__pb2.Robots.SerializeToString,
            robots__pb2.Robots.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetGraph(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/GetGraph',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            groups__pb2.Groups.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/Get',
            robots__pb2.Robot.SerializeToString,
            robots__pb2.Robot.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetDirectusToken(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/GetDirectusToken',
            robots__pb2.Robot.SerializeToString,
            robots__pb2.DirectusToken.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetMyDirectusToken(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/GetMyDirectusToken',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            robots__pb2.DirectusToken.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetByGroup(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/GetByGroup',
            groups__pb2.Group.SerializeToString,
            robots__pb2.Robots.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetSAFile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/GetSAFile',
            robots__pb2.Robot.SerializeToString,
            robots__pb2.SaFile.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Add(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/Add',
            robots__pb2.Robot.SerializeToString,
            robots__pb2.Robot.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SendToRobot(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/SendToRobot',
            robots__pb2.RobotMessage.SerializeToString,
            robots__pb2.RobotMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SendToRobotWithEmail(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/SendToRobotWithEmail',
            robots__pb2.RobotMessage.SerializeToString,
            robots__pb2.RobotMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/Update',
            robots__pb2.Robot.SerializeToString,
            robots__pb2.Robot.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateMacAddress(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/robots.robots/UpdateMacAddress',
            robots__pb2.Robot.SerializeToString,
            robots__pb2.Robot.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
