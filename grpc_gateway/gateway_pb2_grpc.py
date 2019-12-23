# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import gateway_pb2 as gateway__pb2


class GatewayStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ActivateJobs = channel.unary_stream(
        '/gateway_protocol.Gateway/ActivateJobs',
        request_serializer=gateway__pb2.ActivateJobsRequest.SerializeToString,
        response_deserializer=gateway__pb2.ActivateJobsResponse.FromString,
        )
    self.CancelWorkflowInstance = channel.unary_unary(
        '/gateway_protocol.Gateway/CancelWorkflowInstance',
        request_serializer=gateway__pb2.CancelWorkflowInstanceRequest.SerializeToString,
        response_deserializer=gateway__pb2.CancelWorkflowInstanceResponse.FromString,
        )
    self.CompleteJob = channel.unary_unary(
        '/gateway_protocol.Gateway/CompleteJob',
        request_serializer=gateway__pb2.CompleteJobRequest.SerializeToString,
        response_deserializer=gateway__pb2.CompleteJobResponse.FromString,
        )
    self.CreateWorkflowInstance = channel.unary_unary(
        '/gateway_protocol.Gateway/CreateWorkflowInstance',
        request_serializer=gateway__pb2.CreateWorkflowInstanceRequest.SerializeToString,
        response_deserializer=gateway__pb2.CreateWorkflowInstanceResponse.FromString,
        )
    self.DeployWorkflow = channel.unary_unary(
        '/gateway_protocol.Gateway/DeployWorkflow',
        request_serializer=gateway__pb2.DeployWorkflowRequest.SerializeToString,
        response_deserializer=gateway__pb2.DeployWorkflowResponse.FromString,
        )
    self.FailJob = channel.unary_unary(
        '/gateway_protocol.Gateway/FailJob',
        request_serializer=gateway__pb2.FailJobRequest.SerializeToString,
        response_deserializer=gateway__pb2.FailJobResponse.FromString,
        )
    self.PublishMessage = channel.unary_unary(
        '/gateway_protocol.Gateway/PublishMessage',
        request_serializer=gateway__pb2.PublishMessageRequest.SerializeToString,
        response_deserializer=gateway__pb2.PublishMessageResponse.FromString,
        )
    self.ResolveIncident = channel.unary_unary(
        '/gateway_protocol.Gateway/ResolveIncident',
        request_serializer=gateway__pb2.ResolveIncidentRequest.SerializeToString,
        response_deserializer=gateway__pb2.ResolveIncidentResponse.FromString,
        )
    self.SetVariables = channel.unary_unary(
        '/gateway_protocol.Gateway/SetVariables',
        request_serializer=gateway__pb2.SetVariablesRequest.SerializeToString,
        response_deserializer=gateway__pb2.SetVariablesResponse.FromString,
        )
    self.Topology = channel.unary_unary(
        '/gateway_protocol.Gateway/Topology',
        request_serializer=gateway__pb2.TopologyRequest.SerializeToString,
        response_deserializer=gateway__pb2.TopologyResponse.FromString,
        )
    self.UpdateJobRetries = channel.unary_unary(
        '/gateway_protocol.Gateway/UpdateJobRetries',
        request_serializer=gateway__pb2.UpdateJobRetriesRequest.SerializeToString,
        response_deserializer=gateway__pb2.UpdateJobRetriesResponse.FromString,
        )


class GatewayServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def ActivateJobs(self, request, context):
    """
    Iterates through all known partitions round-robin and activates up to the requested
    maximum and streams them back to the client as they are activated.
    Errors:
    INVALID_ARGUMENT:
    - type is blank (empty string, null)
    - worker is blank (empty string, null)
    - timeout less than 1
    - maxJobsToActivate is less than 1
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CancelWorkflowInstance(self, request, context):
    """
    Cancels a running workflow instance
    Errors:
    NOT_FOUND:
    - no workflow instance exists with the given key
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CompleteJob(self, request, context):
    """
    Completes a job with the given variables, which allows completing the associated service task.
    Errors:
    NOT_FOUND:
    - no job exists with the given job key. Note that since jobs are removed once completed,
    it could be that this job did exist at some point.
    FAILED_PRECONDITION:
    - the job was marked as failed. In that case, the related incident must be resolved before
    the job can be activated again and completed.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateWorkflowInstance(self, request, context):
    """
    Creates and starts an instance of the specified workflow. The workflow definition to use to
    create the instance can be specified either using its unique key (as returned by
    DeployWorkflow), or using the BPMN process ID and a version. Pass -1 as the version to use the
    latest deployed version. Note that only workflows with none start events can be started through
    this command.
    Errors:
    NOT_FOUND:
    - no workflow with the given key exists (if workflowKey was given)
    - no workflow with the given process ID exists (if bpmnProcessId was given but version was -1)
    - no workflow with the given process ID and version exists (if both bpmnProcessId and version were given)
    FAILED_PRECONDITION:
    - the workflow definition does not contain a none start event; only workflows with none
    start event can be started manually.
    INVALID_ARGUMENT:
    - the given variables argument is not a valid JSON document; it is expected to be a valid
    JSON document where the root node is an object.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeployWorkflow(self, request, context):
    """
    Deploys one or more workflows to Zeebe. Note that this is an atomic call,
    i.e. either all workflows are deployed, or none of them are.
    Errors:
    INVALID_ARGUMENT:
    - no resources given.
    - if at least one resource is invalid. A resource is considered invalid if:
    - it is not a BPMN or YAML file (currently detected through the file extension)
    - the resource data is not deserializable (e.g. detected as BPMN, but it's broken XML)
    - the workflow is invalid (e.g. an event-based gateway has an outgoing sequence flow to a task)
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def FailJob(self, request, context):
    """
    Marks the job as failed; if the retries argument is positive, then the job will be immediately
    activatable again, and a worker could try again to process it. If it is zero or negative however,
    an incident will be raised, tagged with the given errorMessage, and the job will not be
    activatable until the incident is resolved.
    Errors:
    NOT_FOUND:
    - no job was found with the given key
    FAILED_PRECONDITION:
    - the job was not activated
    - the job is already in a failed state, i.e. ran out of retries
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PublishMessage(self, request, context):
    """
    Publishes a single message. Messages are published to specific partitions computed from their
    correlation keys.
    Errors:
    ALREADY_EXISTS:
    - a message with the same ID was previously published (and is still alive)
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ResolveIncident(self, request, context):
    """
    Resolves a given incident. This simply marks the incident as resolved; most likely a call to
    UpdateJobRetries or SetVariables will be necessary to actually resolve the
    problem, following by this call.
    Errors:
    NOT_FOUND:
    - no incident with the given key exists
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetVariables(self, request, context):
    """
    Updates all the variables of a particular scope (e.g. workflow instance, flow element instance)
    from the given JSON document.
    Errors:
    NOT_FOUND:
    - no element with the given elementInstanceKey exists
    INVALID_ARGUMENT:
    - the given variables document is not a valid JSON document; valid documents are expected to
    be JSON documents where the root node is an object.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Topology(self, request, context):
    """
    Obtains the current topology of the cluster the gateway is part of.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateJobRetries(self, request, context):
    """
    Updates the number of retries a job has left. This is mostly useful for jobs that have run out of
    retries, should the underlying problem be solved.
    Errors:
    NOT_FOUND:
    - no job exists with the given key
    INVALID_ARGUMENT:
    - retries is not greater than 0
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_GatewayServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ActivateJobs': grpc.unary_stream_rpc_method_handler(
          servicer.ActivateJobs,
          request_deserializer=gateway__pb2.ActivateJobsRequest.FromString,
          response_serializer=gateway__pb2.ActivateJobsResponse.SerializeToString,
      ),
      'CancelWorkflowInstance': grpc.unary_unary_rpc_method_handler(
          servicer.CancelWorkflowInstance,
          request_deserializer=gateway__pb2.CancelWorkflowInstanceRequest.FromString,
          response_serializer=gateway__pb2.CancelWorkflowInstanceResponse.SerializeToString,
      ),
      'CompleteJob': grpc.unary_unary_rpc_method_handler(
          servicer.CompleteJob,
          request_deserializer=gateway__pb2.CompleteJobRequest.FromString,
          response_serializer=gateway__pb2.CompleteJobResponse.SerializeToString,
      ),
      'CreateWorkflowInstance': grpc.unary_unary_rpc_method_handler(
          servicer.CreateWorkflowInstance,
          request_deserializer=gateway__pb2.CreateWorkflowInstanceRequest.FromString,
          response_serializer=gateway__pb2.CreateWorkflowInstanceResponse.SerializeToString,
      ),
      'DeployWorkflow': grpc.unary_unary_rpc_method_handler(
          servicer.DeployWorkflow,
          request_deserializer=gateway__pb2.DeployWorkflowRequest.FromString,
          response_serializer=gateway__pb2.DeployWorkflowResponse.SerializeToString,
      ),
      'FailJob': grpc.unary_unary_rpc_method_handler(
          servicer.FailJob,
          request_deserializer=gateway__pb2.FailJobRequest.FromString,
          response_serializer=gateway__pb2.FailJobResponse.SerializeToString,
      ),
      'PublishMessage': grpc.unary_unary_rpc_method_handler(
          servicer.PublishMessage,
          request_deserializer=gateway__pb2.PublishMessageRequest.FromString,
          response_serializer=gateway__pb2.PublishMessageResponse.SerializeToString,
      ),
      'ResolveIncident': grpc.unary_unary_rpc_method_handler(
          servicer.ResolveIncident,
          request_deserializer=gateway__pb2.ResolveIncidentRequest.FromString,
          response_serializer=gateway__pb2.ResolveIncidentResponse.SerializeToString,
      ),
      'SetVariables': grpc.unary_unary_rpc_method_handler(
          servicer.SetVariables,
          request_deserializer=gateway__pb2.SetVariablesRequest.FromString,
          response_serializer=gateway__pb2.SetVariablesResponse.SerializeToString,
      ),
      'Topology': grpc.unary_unary_rpc_method_handler(
          servicer.Topology,
          request_deserializer=gateway__pb2.TopologyRequest.FromString,
          response_serializer=gateway__pb2.TopologyResponse.SerializeToString,
      ),
      'UpdateJobRetries': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateJobRetries,
          request_deserializer=gateway__pb2.UpdateJobRetriesRequest.FromString,
          response_serializer=gateway__pb2.UpdateJobRetriesResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'gateway_protocol.Gateway', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))