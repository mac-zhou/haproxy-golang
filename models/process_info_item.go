// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProcessInfoItem process info item
//
// swagger:model process_info_item
type ProcessInfoItem struct {

	// active peers
	ActivePeers *int64 `json:"active_peers,omitempty"`

	// busy polling
	BusyPolling *int64 `json:"busy_polling,omitempty"`

	// bytes out rate
	BytesOutRate *int64 `json:"bytes_out_rate,omitempty"`

	// compress bps in
	CompressBpsIn *int64 `json:"compress_bps_in,omitempty"`

	// compress bps out
	CompressBpsOut *int64 `json:"compress_bps_out,omitempty"`

	// compress bps rate lim
	CompressBpsRateLim *int64 `json:"compress_bps_rate_lim,omitempty"`

	// conn rate
	ConnRate *int64 `json:"conn_rate,omitempty"`

	// conn rate limit
	ConnRateLimit *int64 `json:"conn_rate_limit,omitempty"`

	// connected peers
	ConnectedPeers *int64 `json:"connected_peers,omitempty"`

	// cum conns
	CumConns *int64 `json:"cum_conns,omitempty"`

	// cum req
	CumReq *int64 `json:"cum_req,omitempty"`

	// cum ssl conns
	CumSslConns *int64 `json:"cum_ssl_conns,omitempty"`

	// curr conns
	CurrConns *int64 `json:"curr_conns,omitempty"`

	// curr ssl conns
	CurrSslConns *int64 `json:"curr_ssl_conns,omitempty"`

	// dropped logs
	DroppedLogs *int64 `json:"dropped_logs,omitempty"`

	// failed resolutions
	FailedResolutions *int64 `json:"failed_resolutions,omitempty"`

	// hard max conn
	HardMaxConn *int64 `json:"hard_max_conn,omitempty"`

	// idle pct
	IdlePct *int64 `json:"idle_pct,omitempty"`

	// jobs
	Jobs *int64 `json:"jobs,omitempty"`

	// listeners
	Listeners *int64 `json:"listeners,omitempty"`

	// max conn
	MaxConn *int64 `json:"max_conn,omitempty"`

	// max conn rate
	MaxConnRate *int64 `json:"max_conn_rate,omitempty"`

	// max pipes
	MaxPipes *int64 `json:"max_pipes,omitempty"`

	// max sess rate
	MaxSessRate *int64 `json:"max_sess_rate,omitempty"`

	// max sock
	MaxSock *int64 `json:"max_sock,omitempty"`

	// max ssl conns
	MaxSslConns *int64 `json:"max_ssl_conns,omitempty"`

	// max ssl rate
	MaxSslRate *int64 `json:"max_ssl_rate,omitempty"`

	// max zlib mem usage
	MaxZlibMemUsage *int64 `json:"max_zlib_mem_usage,omitempty"`

	// mem max mb
	MemMaxMb *int64 `json:"mem_max_mb,omitempty"`

	// Number of threads
	Nbthread *int64 `json:"nbthread,omitempty"`

	// node
	Node string `json:"node,omitempty"`

	// Process id of the replying worker process
	Pid *int64 `json:"pid,omitempty"`

	// pipes free
	PipesFree *int64 `json:"pipes_free,omitempty"`

	// pipes used
	PipesUsed *int64 `json:"pipes_used,omitempty"`

	// pool alloc mb
	PoolAllocMb *int64 `json:"pool_alloc_mb,omitempty"`

	// pool failed
	PoolFailed *int64 `json:"pool_failed,omitempty"`

	// pool used mb
	PoolUsedMb *int64 `json:"pool_used_mb,omitempty"`

	// Process number
	ProcessNum *int64 `json:"process_num,omitempty"`

	// Number of spawned processes
	Processes *int64 `json:"processes,omitempty"`

	// HAProxy version release date
	// Format: date
	ReleaseDate strfmt.Date `json:"release_date,omitempty"`

	// run queue
	RunQueue *int64 `json:"run_queue,omitempty"`

	// sess rate
	SessRate *int64 `json:"sess_rate,omitempty"`

	// sess rate limit
	SessRateLimit *int64 `json:"sess_rate_limit,omitempty"`

	// ssl backend key rate
	SslBackendKeyRate *int64 `json:"ssl_backend_key_rate,omitempty"`

	// ssl backend max key rate
	SslBackendMaxKeyRate *int64 `json:"ssl_backend_max_key_rate,omitempty"`

	// ssl cache lookups
	SslCacheLookups *int64 `json:"ssl_cache_lookups,omitempty"`

	// ssl cache misses
	SslCacheMisses *int64 `json:"ssl_cache_misses,omitempty"`

	// ssl frontend key rate
	SslFrontendKeyRate *int64 `json:"ssl_frontend_key_rate,omitempty"`

	// ssl frontend max key rate
	SslFrontendMaxKeyRate *int64 `json:"ssl_frontend_max_key_rate,omitempty"`

	// ssl frontend session reuse
	SslFrontendSessionReuse *int64 `json:"ssl_frontend_session_reuse,omitempty"`

	// ssl rate
	SslRate *int64 `json:"ssl_rate,omitempty"`

	// ssl rate limit
	SslRateLimit *int64 `json:"ssl_rate_limit,omitempty"`

	// stopping
	Stopping *int64 `json:"stopping,omitempty"`

	// tasks
	Tasks *int64 `json:"tasks,omitempty"`

	// total bytes out
	TotalBytesOut *int64 `json:"total_bytes_out,omitempty"`

	// ulimit n
	Ulimitn *int64 `json:"ulimit_n,omitempty"`

	// unstoppable
	Unstoppable *int64 `json:"unstoppable,omitempty"`

	// HAProxy uptime in s
	Uptime *int64 `json:"uptime,omitempty"`

	// HAProxy version string
	Version string `json:"version,omitempty"`

	// zlib mem usage
	ZlibMemUsage *int64 `json:"zlib_mem_usage,omitempty"`
}

// Validate validates this process info item
func (m *ProcessInfoItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessInfoItem) validateReleaseDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("release_date", "body", "date", m.ReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessInfoItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessInfoItem) UnmarshalBinary(b []byte) error {
	var res ProcessInfoItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
