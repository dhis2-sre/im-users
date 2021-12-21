// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Certificate A Certificate represents an X.509 certificate.
//
// swagger:model Certificate
type Certificate struct {

	// authority key Id
	AuthorityKeyID []uint8 `json:"AuthorityKeyId"`

	// BasicConstraintsValid indicates whether IsCA, MaxPathLen,
	// and MaxPathLenZero are valid.
	BasicConstraintsValid bool `json:"BasicConstraintsValid,omitempty"`

	// CRL Distribution Points
	CRLDistributionPoints []string `json:"CRLDistributionPoints"`

	// Subject Alternate Name values. (Note that these values may not be valid
	// if invalid values were contained within a parsed certificate. For
	// example, an element of DNSNames may not be a valid DNS domain name.)
	DNSNames []string `json:"DNSNames"`

	// email addresses
	EmailAddresses []string `json:"EmailAddresses"`

	// excluded DNS domains
	ExcludedDNSDomains []string `json:"ExcludedDNSDomains"`

	// excluded email addresses
	ExcludedEmailAddresses []string `json:"ExcludedEmailAddresses"`

	// excluded IP ranges
	ExcludedIPRanges []*IPNet `json:"ExcludedIPRanges"`

	// excluded URI domains
	ExcludedURIDomains []string `json:"ExcludedURIDomains"`

	// ext key usage
	ExtKeyUsage []ExtKeyUsage `json:"ExtKeyUsage"`

	// Extensions contains raw X.509 extensions. When parsing certificates,
	// this can be used to extract non-critical extensions that are not
	// parsed by this package. When marshaling certificates, the Extensions
	// field is ignored, see ExtraExtensions.
	Extensions []*Extension `json:"Extensions"`

	// ExtraExtensions contains extensions to be copied, raw, into any
	// marshaled certificates. Values override any extensions that would
	// otherwise be produced based on the other fields. The ExtraExtensions
	// field is not populated when parsing certificates, see Extensions.
	ExtraExtensions []*Extension `json:"ExtraExtensions"`

	// IP addresses
	IPAddresses []IP `json:"IPAddresses"`

	// is c a
	IsCA bool `json:"IsCA,omitempty"`

	// issuer
	Issuer *Name `json:"Issuer,omitempty"`

	// issuing certificate URL
	IssuingCertificateURL []string `json:"IssuingCertificateURL"`

	// key usage
	KeyUsage KeyUsage `json:"KeyUsage,omitempty"`

	// MaxPathLen and MaxPathLenZero indicate the presence and
	// value of the BasicConstraints' "pathLenConstraint".
	//
	// When parsing a certificate, a positive non-zero MaxPathLen
	// means that the field was specified, -1 means it was unset,
	// and MaxPathLenZero being true mean that the field was
	// explicitly set to zero. The case of MaxPathLen==0 with MaxPathLenZero==false
	// should be treated equivalent to -1 (unset).
	//
	// When generating a certificate, an unset pathLenConstraint
	// can be requested with either MaxPathLen == -1 or using the
	// zero value for both MaxPathLen and MaxPathLenZero.
	MaxPathLen int64 `json:"MaxPathLen,omitempty"`

	// MaxPathLenZero indicates that BasicConstraintsValid==true
	// and MaxPathLen==0 should be interpreted as an actual
	// maximum path length of zero. Otherwise, that combination is
	// interpreted as MaxPathLen not being set.
	MaxPathLenZero bool `json:"MaxPathLenZero,omitempty"`

	// not after
	// Format: date-time
	NotAfter strfmt.DateTime `json:"NotBefore,omitempty"`

	// RFC 5280, 4.2.2.1 (Authority Information Access)
	OCSPServer []string `json:"OCSPServer"`

	// permitted DNS domains
	PermittedDNSDomains []string `json:"PermittedDNSDomains"`

	// Name constraints
	PermittedDNSDomainsCritical bool `json:"PermittedDNSDomainsCritical,omitempty"`

	// permitted email addresses
	PermittedEmailAddresses []string `json:"PermittedEmailAddresses"`

	// permitted IP ranges
	PermittedIPRanges []*IPNet `json:"PermittedIPRanges"`

	// permitted URI domains
	PermittedURIDomains []string `json:"PermittedURIDomains"`

	// policy identifiers
	PolicyIdentifiers []ObjectIdentifier `json:"PolicyIdentifiers"`

	// public key
	PublicKey interface{} `json:"PublicKey,omitempty"`

	// public key algorithm
	PublicKeyAlgorithm PublicKeyAlgorithm `json:"PublicKeyAlgorithm,omitempty"`

	// raw
	Raw []uint8 `json:"Raw"`

	// raw issuer
	RawIssuer []uint8 `json:"RawIssuer"`

	// raw subject
	RawSubject []uint8 `json:"RawSubject"`

	// raw subject public key info
	RawSubjectPublicKeyInfo []uint8 `json:"RawSubjectPublicKeyInfo"`

	// raw t b s certificate
	RawTBSCertificate []uint8 `json:"RawTBSCertificate"`

	// serial number
	SerialNumber Int `json:"SerialNumber,omitempty"`

	// signature
	Signature []uint8 `json:"Signature"`

	// signature algorithm
	SignatureAlgorithm SignatureAlgorithm `json:"SignatureAlgorithm,omitempty"`

	// subject
	Subject *Name `json:"Subject,omitempty"`

	// subject key Id
	SubjectKeyID []uint8 `json:"SubjectKeyId"`

	// u r is
	URIs []*URL `json:"URIs"`

	// UnhandledCriticalExtensions contains a list of extension IDs that
	// were not (fully) processed when parsing. Verify will fail if this
	// slice is non-empty, unless verification is delegated to an OS
	// library which understands all the critical extensions.
	//
	// Users can access these extensions using Extensions and can remove
	// elements from this slice if they believe that they have been
	// handled.
	UnhandledCriticalExtensions []ObjectIdentifier `json:"UnhandledCriticalExtensions"`

	// unknown ext key usage
	UnknownExtKeyUsage []ObjectIdentifier `json:"UnknownExtKeyUsage"`

	// version
	Version int64 `json:"Version,omitempty"`
}

// Validate validates this certificate
func (m *Certificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExcludedIPRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtKeyUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtraExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermittedIPRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKeyAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatureAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURIs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnhandledCriticalExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnknownExtKeyUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Certificate) validateExcludedIPRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludedIPRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.ExcludedIPRanges); i++ {
		if swag.IsZero(m.ExcludedIPRanges[i]) { // not required
			continue
		}

		if m.ExcludedIPRanges[i] != nil {
			if err := m.ExcludedIPRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExcludedIPRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExcludedIPRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Certificate) validateExtKeyUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtKeyUsage) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtKeyUsage); i++ {

		if err := m.ExtKeyUsage[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ExtKeyUsage" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ExtKeyUsage" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Certificate) validateExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	for i := 0; i < len(m.Extensions); i++ {
		if swag.IsZero(m.Extensions[i]) { // not required
			continue
		}

		if m.Extensions[i] != nil {
			if err := m.Extensions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Extensions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Certificate) validateExtraExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtraExtensions) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtraExtensions); i++ {
		if swag.IsZero(m.ExtraExtensions[i]) { // not required
			continue
		}

		if m.ExtraExtensions[i] != nil {
			if err := m.ExtraExtensions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExtraExtensions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExtraExtensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Certificate) validateIPAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAddresses) { // not required
		return nil
	}

	for i := 0; i < len(m.IPAddresses); i++ {

		if err := m.IPAddresses[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IPAddresses" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("IPAddresses" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Certificate) validateIssuer(formats strfmt.Registry) error {
	if swag.IsZero(m.Issuer) { // not required
		return nil
	}

	if m.Issuer != nil {
		if err := m.Issuer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Issuer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Issuer")
			}
			return err
		}
	}

	return nil
}

func (m *Certificate) validateKeyUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyUsage) { // not required
		return nil
	}

	if err := m.KeyUsage.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("KeyUsage")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("KeyUsage")
		}
		return err
	}

	return nil
}

func (m *Certificate) validateNotAfter(formats strfmt.Registry) error {
	if swag.IsZero(m.NotAfter) { // not required
		return nil
	}

	if err := validate.FormatOf("NotBefore", "body", "date-time", m.NotAfter.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Certificate) validatePermittedIPRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.PermittedIPRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.PermittedIPRanges); i++ {
		if swag.IsZero(m.PermittedIPRanges[i]) { // not required
			continue
		}

		if m.PermittedIPRanges[i] != nil {
			if err := m.PermittedIPRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PermittedIPRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PermittedIPRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Certificate) validatePolicyIdentifiers(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyIdentifiers) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyIdentifiers); i++ {

		if err := m.PolicyIdentifiers[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PolicyIdentifiers" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PolicyIdentifiers" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Certificate) validatePublicKeyAlgorithm(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicKeyAlgorithm) { // not required
		return nil
	}

	if err := m.PublicKeyAlgorithm.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PublicKeyAlgorithm")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PublicKeyAlgorithm")
		}
		return err
	}

	return nil
}

func (m *Certificate) validateSignatureAlgorithm(formats strfmt.Registry) error {
	if swag.IsZero(m.SignatureAlgorithm) { // not required
		return nil
	}

	if err := m.SignatureAlgorithm.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SignatureAlgorithm")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SignatureAlgorithm")
		}
		return err
	}

	return nil
}

func (m *Certificate) validateSubject(formats strfmt.Registry) error {
	if swag.IsZero(m.Subject) { // not required
		return nil
	}

	if m.Subject != nil {
		if err := m.Subject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Subject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Subject")
			}
			return err
		}
	}

	return nil
}

func (m *Certificate) validateURIs(formats strfmt.Registry) error {
	if swag.IsZero(m.URIs) { // not required
		return nil
	}

	for i := 0; i < len(m.URIs); i++ {
		if swag.IsZero(m.URIs[i]) { // not required
			continue
		}

		if m.URIs[i] != nil {
			if err := m.URIs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("URIs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("URIs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Certificate) validateUnhandledCriticalExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.UnhandledCriticalExtensions) { // not required
		return nil
	}

	for i := 0; i < len(m.UnhandledCriticalExtensions); i++ {

		if err := m.UnhandledCriticalExtensions[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UnhandledCriticalExtensions" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UnhandledCriticalExtensions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Certificate) validateUnknownExtKeyUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.UnknownExtKeyUsage) { // not required
		return nil
	}

	for i := 0; i < len(m.UnknownExtKeyUsage); i++ {

		if err := m.UnknownExtKeyUsage[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UnknownExtKeyUsage" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UnknownExtKeyUsage" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this certificate based on the context it is used
func (m *Certificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExcludedIPRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtKeyUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtraExtensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIssuer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePermittedIPRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyIdentifiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicKeyAlgorithm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignatureAlgorithm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURIs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnhandledCriticalExtensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnknownExtKeyUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Certificate) contextValidateExcludedIPRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExcludedIPRanges); i++ {

		if m.ExcludedIPRanges[i] != nil {
			if err := m.ExcludedIPRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExcludedIPRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExcludedIPRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Certificate) contextValidateExtKeyUsage(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExtKeyUsage); i++ {

		if err := m.ExtKeyUsage[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ExtKeyUsage" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ExtKeyUsage" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Certificate) contextValidateExtensions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Extensions); i++ {

		if m.Extensions[i] != nil {
			if err := m.Extensions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Extensions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Certificate) contextValidateExtraExtensions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExtraExtensions); i++ {

		if m.ExtraExtensions[i] != nil {
			if err := m.ExtraExtensions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExtraExtensions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExtraExtensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Certificate) contextValidateIPAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPAddresses); i++ {

		if err := m.IPAddresses[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IPAddresses" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("IPAddresses" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Certificate) contextValidateIssuer(ctx context.Context, formats strfmt.Registry) error {

	if m.Issuer != nil {
		if err := m.Issuer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Issuer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Issuer")
			}
			return err
		}
	}

	return nil
}

func (m *Certificate) contextValidateKeyUsage(ctx context.Context, formats strfmt.Registry) error {

	if err := m.KeyUsage.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("KeyUsage")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("KeyUsage")
		}
		return err
	}

	return nil
}

func (m *Certificate) contextValidatePermittedIPRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PermittedIPRanges); i++ {

		if m.PermittedIPRanges[i] != nil {
			if err := m.PermittedIPRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PermittedIPRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PermittedIPRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Certificate) contextValidatePolicyIdentifiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PolicyIdentifiers); i++ {

		if err := m.PolicyIdentifiers[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PolicyIdentifiers" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PolicyIdentifiers" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Certificate) contextValidatePublicKeyAlgorithm(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PublicKeyAlgorithm.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PublicKeyAlgorithm")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PublicKeyAlgorithm")
		}
		return err
	}

	return nil
}

func (m *Certificate) contextValidateSignatureAlgorithm(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SignatureAlgorithm.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SignatureAlgorithm")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SignatureAlgorithm")
		}
		return err
	}

	return nil
}

func (m *Certificate) contextValidateSubject(ctx context.Context, formats strfmt.Registry) error {

	if m.Subject != nil {
		if err := m.Subject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Subject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Subject")
			}
			return err
		}
	}

	return nil
}

func (m *Certificate) contextValidateURIs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.URIs); i++ {

		if m.URIs[i] != nil {
			if err := m.URIs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("URIs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("URIs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Certificate) contextValidateUnhandledCriticalExtensions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UnhandledCriticalExtensions); i++ {

		if err := m.UnhandledCriticalExtensions[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UnhandledCriticalExtensions" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UnhandledCriticalExtensions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Certificate) contextValidateUnknownExtKeyUsage(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UnknownExtKeyUsage); i++ {

		if err := m.UnknownExtKeyUsage[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UnknownExtKeyUsage" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UnknownExtKeyUsage" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Certificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Certificate) UnmarshalBinary(b []byte) error {
	var res Certificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
