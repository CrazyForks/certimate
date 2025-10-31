package handlers

import (
	"context"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"

	"github.com/certimate-go/certimate/internal/domain/dtos"
	"github.com/certimate-go/certimate/internal/rest/resp"
)

type certificateService interface {
	DownloadArchivedFile(ctx context.Context, req *dtos.CertificateArchiveFileReq) (*dtos.CertificateArchiveFileResp, error)
	RevokeCertificate(ctx context.Context, req *dtos.CertificateRevokeReq) (*dtos.CertificateRevokeResp, error)
	ValidateCertificate(ctx context.Context, req *dtos.CertificateValidateCertificateReq) (*dtos.CertificateValidateCertificateResp, error)
	ValidatePrivateKey(ctx context.Context, req *dtos.CertificateValidatePrivateKeyReq) (*dtos.CertificateValidatePrivateKeyResp, error)
}

type CertificateHandler struct {
	service certificateService
}

func NewCertificateHandler(router *router.RouterGroup[*core.RequestEvent], service certificateService) {
	handler := &CertificateHandler{
		service: service,
	}

	group := router.Group("/certificates")
	group.POST("/{certificateId}/archive", handler.archiveCertificate)
	group.POST("/{certificateId}/revoke", handler.revokeCertificate)
	group.POST("/validate/certificate", handler.validateCertificate)
	group.POST("/validate/private-key", handler.validatePrivateKey)
}

func (handler *CertificateHandler) archiveCertificate(e *core.RequestEvent) error {
	req := &dtos.CertificateArchiveFileReq{}
	req.CertificateId = e.Request.PathValue("certificateId")
	if err := e.BindBody(req); err != nil {
		return resp.Err(e, err)
	}

	res, err := handler.service.DownloadArchivedFile(e.Request.Context(), req)
	if err != nil {
		return resp.Err(e, err)
	}

	return resp.Ok(e, res)
}

func (handler *CertificateHandler) revokeCertificate(e *core.RequestEvent) error {
	req := &dtos.CertificateRevokeReq{}
	req.CertificateId = e.Request.PathValue("certificateId")
	if err := e.BindBody(req); err != nil {
		return resp.Err(e, err)
	}

	res, err := handler.service.RevokeCertificate(e.Request.Context(), req)
	if err != nil {
		return resp.Err(e, err)
	}

	return resp.Ok(e, res)
}

func (handler *CertificateHandler) validateCertificate(e *core.RequestEvent) error {
	req := &dtos.CertificateValidateCertificateReq{}
	if err := e.BindBody(req); err != nil {
		return resp.Err(e, err)
	}

	res, err := handler.service.ValidateCertificate(e.Request.Context(), req)
	if err != nil {
		return resp.Err(e, err)
	}

	return resp.Ok(e, res)
}

func (handler *CertificateHandler) validatePrivateKey(e *core.RequestEvent) error {
	req := &dtos.CertificateValidatePrivateKeyReq{}
	if err := e.BindBody(req); err != nil {
		return resp.Err(e, err)
	}

	res, err := handler.service.ValidatePrivateKey(e.Request.Context(), req)
	if err != nil {
		return resp.Err(e, err)
	}

	return resp.Ok(e, res)
}
