package trace

type IgnoredTracerContext struct {
	noopSpan NoopSpan
	noopEntrySpan  NoopEntrySpan
	spans stack
}

var _ TracerContext =(*IgnoredTracerContext)(nil)


func (c *IgnoredTracerContext) inject(carrier ContextCarrier){

}
func (c *IgnoredTracerContext) extract(carrier ContextCarrier){

}
func (c *IgnoredTracerContext) capture()  *ContextSnapshot{
	return  &ContextSnapshot{}
}
func (c *IgnoredTracerContext) continued(snapshot ContextSnapshot){

}
func (c *IgnoredTracerContext) getReadableGlobalTraceId() string{
	return ""
}
func (c *IgnoredTracerContext) createEntrySpan(operationName string) Span{
	c.spans.push(c.noopEntrySpan)
	return  &c.noopEntrySpan
}
func (c *IgnoredTracerContext) createLocalSpan(operationName string) Span{
	c.spans.push(c.noopEntrySpan)
	return  &c.noopEntrySpan
}
func (c *IgnoredTracerContext) createExitSpan(operationName string,remotePeer string) Span{
	exitSpan :=NewNoopExitSpan(0,remotePeer)
	c.spans.push(exitSpan)
	return  exitSpan
}
func (c *IgnoredTracerContext) stopSpan(span Span){
	//_spans.TryPop(out _);
	//if (_spans.Count == 0)
	//{
	//	ListenerManager.NotifyFinish(this);
	//	foreach (var item in Properties)
	//	{
	//	if (item.Value is IDisposable disposable)
	//	{
	//	disposable.Dispose();
	//	}
	//	}
	//}
}
