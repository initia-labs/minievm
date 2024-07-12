package backend

// ClientVersion returns the node name
func (b *JSONRPCBackend) ClientVersion() (string, error) {
	status, err := b.clientCtx.Client.Status(b.ctx)
	if err != nil {
		return "", err
	}

	return status.NodeInfo.Version, nil
}
