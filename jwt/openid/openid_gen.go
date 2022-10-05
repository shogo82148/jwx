// Generated by "sketch" utility. DO NOT EDIT
package openid

import (
	"bytes"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/lestrrat-go/blackmagic"
	"github.com/lestrrat-go/jwx/v2/internal/json"
	"github.com/lestrrat-go/jwx/v2/jwt/internal/types"
)

type stdToken struct {
	mu                  sync.RWMutex
	address             *AddressClaim
	audience            types.Audience
	birthdate           *BirthdateClaim
	email               *string
	emailVerified       *bool
	expiration          *types.NumericDate
	familyName          *string
	gender              *string
	givenName           *string
	issuedAt            *types.NumericDate
	issuer              *string
	jwtID               *string
	locale              *string
	middleName          *string
	name                *string
	nickname            *string
	notBefore           *types.NumericDate
	phoneNumber         *string
	phoneNumberVerified *bool
	picture             *string
	preferredUsername   *string
	profile             *string
	subject             *string
	updatedAt           *types.NumericDate
	website             *string
	zoneinfo            *string
	dc                  DecodeCtx
	extra               map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	AddressKey             = "address"
	AudienceKey            = "aud"
	BirthdateKey           = "birthdate"
	EmailKey               = "email"
	EmailVerifiedKey       = "email_verified"
	ExpirationKey          = "exp"
	FamilyNameKey          = "family_name"
	GenderKey              = "gender"
	GivenNameKey           = "given_name"
	IssuedAtKey            = "iat"
	IssuerKey              = "iss"
	JwtIDKey               = "jti"
	LocaleKey              = "locale"
	MiddleNameKey          = "middle_name"
	NameKey                = "name"
	NicknameKey            = "nickname"
	NotBeforeKey           = "nbf"
	PhoneNumberKey         = "phone_number"
	PhoneNumberVerifiedKey = "phone_number_verified"
	PictureKey             = "picture"
	PreferredUsernameKey   = "preferred_username"
	ProfileKey             = "profile"
	SubjectKey             = "sub"
	UpdatedAtKey           = "updated_at"
	WebsiteKey             = "website"
	ZoneinfoKey            = "zoneinfo"
)

// Get retrieves the value associated with a key
func (v *stdToken) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.getNoLock(key, dst, false)
}

// getNoLock is a utility method that is called from Get, MarshalJSON, etc, but
// it can be used from user-supplied code. Unlike Get, it avoids locking for
// each call, so the user needs to explicitly lock the object before using,
// but otherwise should be faster than sing Get directly
func (v *stdToken) getNoLock(key string, dst interface{}, raw bool) error {
	switch key {
	case AddressKey:
		if val := v.address; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case AudienceKey:
		if val := v.audience; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.GetValue())
		}
	case BirthdateKey:
		if val := v.birthdate; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case EmailKey:
		if val := v.email; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case EmailVerifiedKey:
		if val := v.emailVerified; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ExpirationKey:
		if val := v.expiration; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.GetValue())
		}
	case FamilyNameKey:
		if val := v.familyName; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case GenderKey:
		if val := v.gender; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case GivenNameKey:
		if val := v.givenName; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case IssuedAtKey:
		if val := v.issuedAt; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.GetValue())
		}
	case IssuerKey:
		if val := v.issuer; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case JwtIDKey:
		if val := v.jwtID; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case LocaleKey:
		if val := v.locale; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case MiddleNameKey:
		if val := v.middleName; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case NameKey:
		if val := v.name; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case NicknameKey:
		if val := v.nickname; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case NotBeforeKey:
		if val := v.notBefore; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.GetValue())
		}
	case PhoneNumberKey:
		if val := v.phoneNumber; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case PhoneNumberVerifiedKey:
		if val := v.phoneNumberVerified; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case PictureKey:
		if val := v.picture; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case PreferredUsernameKey:
		if val := v.preferredUsername; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ProfileKey:
		if val := v.profile; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case SubjectKey:
		if val := v.subject; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case UpdatedAtKey:
		if val := v.updatedAt; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.GetValue())
		}
	case WebsiteKey:
		if val := v.website; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case ZoneinfoKey:
		if val := v.zoneinfo; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	default:
		if v.extra != nil {
			val, ok := v.extra[key]
			if ok {
				return blackmagic.AssignIfCompatible(dst, val)
			}
		}
	}
	return fmt.Errorf(`no such key %q`, key)
}

// Set sets the value of the specified field. The name must be a JSON
// field name, not the Go name
func (v *stdToken) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case AddressKey:
		var object AddressClaim
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.address = &object
	case AudienceKey:
		var object types.Audience
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.audience = object
	case BirthdateKey:
		var object BirthdateClaim
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.birthdate = &object
	case EmailKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field email, got %T`, value)
		}
		v.email = &converted
	case EmailVerifiedKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field email_verified, got %T`, value)
		}
		v.emailVerified = &converted
	case ExpirationKey:
		var object types.NumericDate
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.expiration = &object
	case FamilyNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field family_name, got %T`, value)
		}
		v.familyName = &converted
	case GenderKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field gender, got %T`, value)
		}
		v.gender = &converted
	case GivenNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field given_name, got %T`, value)
		}
		v.givenName = &converted
	case IssuedAtKey:
		var object types.NumericDate
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.issuedAt = &object
	case IssuerKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field iss, got %T`, value)
		}
		v.issuer = &converted
	case JwtIDKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field jti, got %T`, value)
		}
		v.jwtID = &converted
	case LocaleKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field locale, got %T`, value)
		}
		v.locale = &converted
	case MiddleNameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field middle_name, got %T`, value)
		}
		v.middleName = &converted
	case NameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field name, got %T`, value)
		}
		v.name = &converted
	case NicknameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field nickname, got %T`, value)
		}
		v.nickname = &converted
	case NotBeforeKey:
		var object types.NumericDate
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.notBefore = &object
	case PhoneNumberKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field phone_number, got %T`, value)
		}
		v.phoneNumber = &converted
	case PhoneNumberVerifiedKey:
		converted, ok := value.(bool)
		if !ok {
			return fmt.Errorf(`expected value of type bool for field phone_number_verified, got %T`, value)
		}
		v.phoneNumberVerified = &converted
	case PictureKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field picture, got %T`, value)
		}
		v.picture = &converted
	case PreferredUsernameKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field preferred_username, got %T`, value)
		}
		v.preferredUsername = &converted
	case ProfileKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field profile, got %T`, value)
		}
		v.profile = &converted
	case SubjectKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field sub, got %T`, value)
		}
		v.subject = &converted
	case UpdatedAtKey:
		var object types.NumericDate
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.updatedAt = &object
	case WebsiteKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field website, got %T`, value)
		}
		v.website = &converted
	case ZoneinfoKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field zoneinfo, got %T`, value)
		}
		v.zoneinfo = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}

// Has returns true if the field specified by the argument has been populated.
// The field name must be the JSON field name, not the Go-structure's field name.
func (v *stdToken) Has(name string) bool {
	switch name {
	case AddressKey:
		return v.address != nil
	case AudienceKey:
		return v.audience != nil
	case BirthdateKey:
		return v.birthdate != nil
	case EmailKey:
		return v.email != nil
	case EmailVerifiedKey:
		return v.emailVerified != nil
	case ExpirationKey:
		return v.expiration != nil
	case FamilyNameKey:
		return v.familyName != nil
	case GenderKey:
		return v.gender != nil
	case GivenNameKey:
		return v.givenName != nil
	case IssuedAtKey:
		return v.issuedAt != nil
	case IssuerKey:
		return v.issuer != nil
	case JwtIDKey:
		return v.jwtID != nil
	case LocaleKey:
		return v.locale != nil
	case MiddleNameKey:
		return v.middleName != nil
	case NameKey:
		return v.name != nil
	case NicknameKey:
		return v.nickname != nil
	case NotBeforeKey:
		return v.notBefore != nil
	case PhoneNumberKey:
		return v.phoneNumber != nil
	case PhoneNumberVerifiedKey:
		return v.phoneNumberVerified != nil
	case PictureKey:
		return v.picture != nil
	case PreferredUsernameKey:
		return v.preferredUsername != nil
	case ProfileKey:
		return v.profile != nil
	case SubjectKey:
		return v.subject != nil
	case UpdatedAtKey:
		return v.updatedAt != nil
	case WebsiteKey:
		return v.website != nil
	case ZoneinfoKey:
		return v.zoneinfo != nil
	default:
		if v.extra != nil {
			if _, ok := v.extra[name]; ok {
				return true
			}
		}
		return false
	}
}

// Keys returns a slice of string comprising of JSON field names whose values
// are present in the object.
func (v *stdToken) Keys() []string {
	keys := make([]string, 0, 27)
	if v.address != nil {
		keys = append(keys, AddressKey)
	}
	if v.audience != nil {
		keys = append(keys, AudienceKey)
	}
	if v.birthdate != nil {
		keys = append(keys, BirthdateKey)
	}
	if v.email != nil {
		keys = append(keys, EmailKey)
	}
	if v.emailVerified != nil {
		keys = append(keys, EmailVerifiedKey)
	}
	if v.expiration != nil {
		keys = append(keys, ExpirationKey)
	}
	if v.familyName != nil {
		keys = append(keys, FamilyNameKey)
	}
	if v.gender != nil {
		keys = append(keys, GenderKey)
	}
	if v.givenName != nil {
		keys = append(keys, GivenNameKey)
	}
	if v.issuedAt != nil {
		keys = append(keys, IssuedAtKey)
	}
	if v.issuer != nil {
		keys = append(keys, IssuerKey)
	}
	if v.jwtID != nil {
		keys = append(keys, JwtIDKey)
	}
	if v.locale != nil {
		keys = append(keys, LocaleKey)
	}
	if v.middleName != nil {
		keys = append(keys, MiddleNameKey)
	}
	if v.name != nil {
		keys = append(keys, NameKey)
	}
	if v.nickname != nil {
		keys = append(keys, NicknameKey)
	}
	if v.notBefore != nil {
		keys = append(keys, NotBeforeKey)
	}
	if v.phoneNumber != nil {
		keys = append(keys, PhoneNumberKey)
	}
	if v.phoneNumberVerified != nil {
		keys = append(keys, PhoneNumberVerifiedKey)
	}
	if v.picture != nil {
		keys = append(keys, PictureKey)
	}
	if v.preferredUsername != nil {
		keys = append(keys, PreferredUsernameKey)
	}
	if v.profile != nil {
		keys = append(keys, ProfileKey)
	}
	if v.subject != nil {
		keys = append(keys, SubjectKey)
	}
	if v.updatedAt != nil {
		keys = append(keys, UpdatedAtKey)
	}
	if v.website != nil {
		keys = append(keys, WebsiteKey)
	}
	if v.zoneinfo != nil {
		keys = append(keys, ZoneinfoKey)
	}

	if len(v.extra) > 0 {
		for k := range v.extra {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// HasAddress returns true if the field `address` has been populated
func (v *stdToken) HasAddress() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.address != nil
}

// HasAudience returns true if the field `aud` has been populated
func (v *stdToken) HasAudience() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.audience != nil
}

// HasBirthdate returns true if the field `birthdate` has been populated
func (v *stdToken) HasBirthdate() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.birthdate != nil
}

// HasEmail returns true if the field `email` has been populated
func (v *stdToken) HasEmail() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.email != nil
}

// HasEmailVerified returns true if the field `email_verified` has been populated
func (v *stdToken) HasEmailVerified() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.emailVerified != nil
}

// HasExpiration returns true if the field `exp` has been populated
func (v *stdToken) HasExpiration() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.expiration != nil
}

// HasFamilyName returns true if the field `family_name` has been populated
func (v *stdToken) HasFamilyName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.familyName != nil
}

// HasGender returns true if the field `gender` has been populated
func (v *stdToken) HasGender() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.gender != nil
}

// HasGivenName returns true if the field `given_name` has been populated
func (v *stdToken) HasGivenName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.givenName != nil
}

// HasIssuedAt returns true if the field `iat` has been populated
func (v *stdToken) HasIssuedAt() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.issuedAt != nil
}

// HasIssuer returns true if the field `iss` has been populated
func (v *stdToken) HasIssuer() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.issuer != nil
}

// HasJwtID returns true if the field `jti` has been populated
func (v *stdToken) HasJwtID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.jwtID != nil
}

// HasLocale returns true if the field `locale` has been populated
func (v *stdToken) HasLocale() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.locale != nil
}

// HasMiddleName returns true if the field `middle_name` has been populated
func (v *stdToken) HasMiddleName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.middleName != nil
}

// HasName returns true if the field `name` has been populated
func (v *stdToken) HasName() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.name != nil
}

// HasNickname returns true if the field `nickname` has been populated
func (v *stdToken) HasNickname() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.nickname != nil
}

// HasNotBefore returns true if the field `nbf` has been populated
func (v *stdToken) HasNotBefore() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.notBefore != nil
}

// HasPhoneNumber returns true if the field `phone_number` has been populated
func (v *stdToken) HasPhoneNumber() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.phoneNumber != nil
}

// HasPhoneNumberVerified returns true if the field `phone_number_verified` has been populated
func (v *stdToken) HasPhoneNumberVerified() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.phoneNumberVerified != nil
}

// HasPicture returns true if the field `picture` has been populated
func (v *stdToken) HasPicture() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.picture != nil
}

// HasPreferredUsername returns true if the field `preferred_username` has been populated
func (v *stdToken) HasPreferredUsername() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.preferredUsername != nil
}

// HasProfile returns true if the field `profile` has been populated
func (v *stdToken) HasProfile() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.profile != nil
}

// HasSubject returns true if the field `sub` has been populated
func (v *stdToken) HasSubject() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.subject != nil
}

// HasUpdatedAt returns true if the field `updated_at` has been populated
func (v *stdToken) HasUpdatedAt() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.updatedAt != nil
}

// HasWebsite returns true if the field `website` has been populated
func (v *stdToken) HasWebsite() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.website != nil
}

// HasZoneinfo returns true if the field `zoneinfo` has been populated
func (v *stdToken) HasZoneinfo() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.zoneinfo != nil
}

func (v *stdToken) Address() *AddressClaim {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.address; val != nil {
		return val
	}
	return nil
}

func (v *stdToken) Audience() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.audience; val != nil {
		return val.GetValue()
	}
	return nil
}

func (v *stdToken) Birthdate() *BirthdateClaim {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.birthdate; val != nil {
		return val
	}
	return nil
}

func (v *stdToken) Email() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.email; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) EmailVerified() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.emailVerified; val != nil {
		return *val
	}
	return false
}

func (v *stdToken) Expiration() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.expiration; val != nil {
		return val.GetValue()
	}
	return time.Time{}
}

func (v *stdToken) FamilyName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.familyName; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) Gender() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.gender; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) GivenName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.givenName; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) IssuedAt() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.issuedAt; val != nil {
		return val.GetValue()
	}
	return time.Time{}
}

func (v *stdToken) Issuer() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.issuer; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) JwtID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.jwtID; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) Locale() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.locale; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) MiddleName() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.middleName; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) Name() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.name; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) Nickname() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.nickname; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) NotBefore() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.notBefore; val != nil {
		return val.GetValue()
	}
	return time.Time{}
}

func (v *stdToken) PhoneNumber() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.phoneNumber; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) PhoneNumberVerified() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.phoneNumberVerified; val != nil {
		return *val
	}
	return false
}

func (v *stdToken) Picture() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.picture; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) PreferredUsername() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.preferredUsername; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) Profile() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.profile; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) Subject() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.subject; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) UpdatedAt() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.updatedAt; val != nil {
		return val.GetValue()
	}
	return time.Time{}
}

func (v *stdToken) Website() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.website; val != nil {
		return *val
	}
	return ""
}

func (v *stdToken) Zoneinfo() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.zoneinfo; val != nil {
		return *val
	}
	return ""
}

// Remove removes the value associated with a key
func (v *stdToken) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case AddressKey:
		v.address = nil
	case AudienceKey:
		v.audience = nil
	case BirthdateKey:
		v.birthdate = nil
	case EmailKey:
		v.email = nil
	case EmailVerifiedKey:
		v.emailVerified = nil
	case ExpirationKey:
		v.expiration = nil
	case FamilyNameKey:
		v.familyName = nil
	case GenderKey:
		v.gender = nil
	case GivenNameKey:
		v.givenName = nil
	case IssuedAtKey:
		v.issuedAt = nil
	case IssuerKey:
		v.issuer = nil
	case JwtIDKey:
		v.jwtID = nil
	case LocaleKey:
		v.locale = nil
	case MiddleNameKey:
		v.middleName = nil
	case NameKey:
		v.name = nil
	case NicknameKey:
		v.nickname = nil
	case NotBeforeKey:
		v.notBefore = nil
	case PhoneNumberKey:
		v.phoneNumber = nil
	case PhoneNumberVerifiedKey:
		v.phoneNumberVerified = nil
	case PictureKey:
		v.picture = nil
	case PreferredUsernameKey:
		v.preferredUsername = nil
	case ProfileKey:
		v.profile = nil
	case SubjectKey:
		v.subject = nil
	case UpdatedAtKey:
		v.updatedAt = nil
	case WebsiteKey:
		v.website = nil
	case ZoneinfoKey:
		v.zoneinfo = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *stdToken) Clone(dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	extra := make(map[string]interface{})
	for key, val := range v.extra {
		extra[key] = val
	}
	return blackmagic.AssignIfCompatible(dst, &stdToken{
		address:             v.address,
		audience:            v.audience,
		birthdate:           v.birthdate,
		email:               v.email,
		emailVerified:       v.emailVerified,
		expiration:          v.expiration,
		familyName:          v.familyName,
		gender:              v.gender,
		givenName:           v.givenName,
		issuedAt:            v.issuedAt,
		issuer:              v.issuer,
		jwtID:               v.jwtID,
		locale:              v.locale,
		middleName:          v.middleName,
		name:                v.name,
		nickname:            v.nickname,
		notBefore:           v.notBefore,
		phoneNumber:         v.phoneNumber,
		phoneNumberVerified: v.phoneNumberVerified,
		picture:             v.picture,
		preferredUsername:   v.preferredUsername,
		profile:             v.profile,
		subject:             v.subject,
		updatedAt:           v.updatedAt,
		website:             v.website,
		zoneinfo:            v.zoneinfo,
		dc:                  v.dc,
		extra:               extra,
	})
}

// MarshalJSON serializes stdToken into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *stdToken) MarshalJSON() ([]byte, error) {
	v.mu.RLock()
	defer v.mu.RUnlock()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	buf.WriteByte('{')
	for i, k := range v.Keys() {
		var val interface{}
		if err := v.getNoLock(k, &val, true); err != nil {
			return nil, fmt.Errorf(`failed to retrieve value for field %q: %w`, k, err)
		}

		if i > 0 {
			buf.WriteByte(',')
		}
		if err := enc.Encode(k); err != nil {
			return nil, fmt.Errorf(`failed to encode map key name: %w`, err)
		}
		buf.WriteByte(':')
		if err := enc.Encode(val); err != nil {
			return nil, fmt.Errorf(`failed to encode map value for %q: %w`, k, err)
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// UnmarshalJSON deserializes a piece of JSON data into stdToken.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *stdToken) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.address = nil
	v.audience = nil
	v.birthdate = nil
	v.email = nil
	v.emailVerified = nil
	v.expiration = nil
	v.familyName = nil
	v.gender = nil
	v.givenName = nil
	v.issuedAt = nil
	v.issuer = nil
	v.jwtID = nil
	v.locale = nil
	v.middleName = nil
	v.name = nil
	v.nickname = nil
	v.notBefore = nil
	v.phoneNumber = nil
	v.phoneNumberVerified = nil
	v.picture = nil
	v.preferredUsername = nil
	v.profile = nil
	v.subject = nil
	v.updatedAt = nil
	v.website = nil
	v.zoneinfo = nil

	dec := json.NewDecoder(bytes.NewReader(data))
	var extra map[string]interface{}

LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf(`error reading JSON token: %w`, err)
		}
		switch tok := tok.(type) {
		case json.Delim:
			if tok == '}' { // end of object
				break LOOP
			}
			// we should only get into this clause at the very beginning, and just once
			if tok != '{' {
				return fmt.Errorf(`expected '{', but got '%c'`, tok)
			}
		case string:
			switch tok {
			case AddressKey:
				var val AddressClaim
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AddressKey, err)
				}
				v.address = &val
			case AudienceKey:
				var val types.Audience
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AudienceKey, err)
				}
				v.audience = val
			case BirthdateKey:
				var val BirthdateClaim
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, BirthdateKey, err)
				}
				v.birthdate = &val
			case EmailKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EmailKey, err)
				}
				v.email = &val
			case EmailVerifiedKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, EmailVerifiedKey, err)
				}
				v.emailVerified = &val
			case ExpirationKey:
				var val types.NumericDate
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ExpirationKey, err)
				}
				v.expiration = &val
			case FamilyNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, FamilyNameKey, err)
				}
				v.familyName = &val
			case GenderKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GenderKey, err)
				}
				v.gender = &val
			case GivenNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, GivenNameKey, err)
				}
				v.givenName = &val
			case IssuedAtKey:
				var val types.NumericDate
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, IssuedAtKey, err)
				}
				v.issuedAt = &val
			case IssuerKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, IssuerKey, err)
				}
				v.issuer = &val
			case JwtIDKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, JwtIDKey, err)
				}
				v.jwtID = &val
			case LocaleKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, LocaleKey, err)
				}
				v.locale = &val
			case MiddleNameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, MiddleNameKey, err)
				}
				v.middleName = &val
			case NameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, NameKey, err)
				}
				v.name = &val
			case NicknameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, NicknameKey, err)
				}
				v.nickname = &val
			case NotBeforeKey:
				var val types.NumericDate
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, NotBeforeKey, err)
				}
				v.notBefore = &val
			case PhoneNumberKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PhoneNumberKey, err)
				}
				v.phoneNumber = &val
			case PhoneNumberVerifiedKey:
				var val bool
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PhoneNumberVerifiedKey, err)
				}
				v.phoneNumberVerified = &val
			case PictureKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PictureKey, err)
				}
				v.picture = &val
			case PreferredUsernameKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, PreferredUsernameKey, err)
				}
				v.preferredUsername = &val
			case ProfileKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ProfileKey, err)
				}
				v.profile = &val
			case SubjectKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SubjectKey, err)
				}
				v.subject = &val
			case UpdatedAtKey:
				var val types.NumericDate
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, UpdatedAtKey, err)
				}
				v.updatedAt = &val
			case WebsiteKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, WebsiteKey, err)
				}
				v.website = &val
			case ZoneinfoKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ZoneinfoKey, err)
				}
				v.zoneinfo = &val
			default:
				var val interface{}
				if err := v.decodeExtraField(tok, dec, &val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, tok, err)
				}
				if extra == nil {
					extra = make(map[string]interface{})
				}
				extra[tok] = val
			}
		}
	}

	if extra != nil {
		v.extra = extra
	}
	return nil
}

type Builder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *stdToken
}

// NewBuilder creates a new Builder instance.
// Builder is safe to be used uninitialized as well.
func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) initialize() {
	b.err = nil
	b.object = &stdToken{}
}
func (b *Builder) Address(in *AddressClaim) *Builder {
	return b.Claim(AddressKey, in)
}
func (b *Builder) Audience(in []string) *Builder {
	return b.Claim(AudienceKey, in)
}
func (b *Builder) Birthdate(in *BirthdateClaim) *Builder {
	return b.Claim(BirthdateKey, in)
}
func (b *Builder) Email(in string) *Builder {
	return b.Claim(EmailKey, in)
}
func (b *Builder) EmailVerified(in bool) *Builder {
	return b.Claim(EmailVerifiedKey, in)
}
func (b *Builder) Expiration(in time.Time) *Builder {
	return b.Claim(ExpirationKey, in)
}
func (b *Builder) FamilyName(in string) *Builder {
	return b.Claim(FamilyNameKey, in)
}
func (b *Builder) Gender(in string) *Builder {
	return b.Claim(GenderKey, in)
}
func (b *Builder) GivenName(in string) *Builder {
	return b.Claim(GivenNameKey, in)
}
func (b *Builder) IssuedAt(in time.Time) *Builder {
	return b.Claim(IssuedAtKey, in)
}
func (b *Builder) Issuer(in string) *Builder {
	return b.Claim(IssuerKey, in)
}
func (b *Builder) JwtID(in string) *Builder {
	return b.Claim(JwtIDKey, in)
}
func (b *Builder) Locale(in string) *Builder {
	return b.Claim(LocaleKey, in)
}
func (b *Builder) MiddleName(in string) *Builder {
	return b.Claim(MiddleNameKey, in)
}
func (b *Builder) Name(in string) *Builder {
	return b.Claim(NameKey, in)
}
func (b *Builder) Nickname(in string) *Builder {
	return b.Claim(NicknameKey, in)
}
func (b *Builder) NotBefore(in time.Time) *Builder {
	return b.Claim(NotBeforeKey, in)
}
func (b *Builder) PhoneNumber(in string) *Builder {
	return b.Claim(PhoneNumberKey, in)
}
func (b *Builder) PhoneNumberVerified(in bool) *Builder {
	return b.Claim(PhoneNumberVerifiedKey, in)
}
func (b *Builder) Picture(in string) *Builder {
	return b.Claim(PictureKey, in)
}
func (b *Builder) PreferredUsername(in string) *Builder {
	return b.Claim(PreferredUsernameKey, in)
}
func (b *Builder) Profile(in string) *Builder {
	return b.Claim(ProfileKey, in)
}
func (b *Builder) Subject(in string) *Builder {
	return b.Claim(SubjectKey, in)
}
func (b *Builder) UpdatedAt(in time.Time) *Builder {
	return b.Claim(UpdatedAtKey, in)
}
func (b *Builder) Website(in string) *Builder {
	return b.Claim(WebsiteKey, in)
}
func (b *Builder) Zoneinfo(in string) *Builder {
	return b.Claim(ZoneinfoKey, in)
}

// Claim sets the value of any field. The name should be the JSON field name.
// Type check will only be performed for pre-defined types
func (b *Builder) Claim(name string, value interface{}) *Builder {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	if err := b.object.Set(name, value); err != nil {
		b.err = err
	}
	return b
}
func (b *Builder) Build() (Token, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return nil, b.err
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}
func (b *Builder) MustBuild() Token {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

// New creates an empty token
func New() Token {
	return &stdToken{}
}

func (v *stdToken) DecodeCtx() DecodeCtx {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.dc
}

func (v *stdToken) SetDecodeCtx(dc DecodeCtx) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.dc = dc
}

func (v *stdToken) decodeExtraField(name string, dec *json.Decoder, dst interface{}) error {
	if dc := v.dc; dc != nil {
		if localReg := dc.Registry(); localReg != nil {
			decoded, err := localReg.Decode(dec, name)
			if err == nil {
				if err := blackmagic.AssignIfCompatible(dst, decoded); err != nil {
					return fmt.Errorf(`failed to assign decoded value for %q: %w`, name, err)
				}
				return nil
			}
		}
	}

	decoded, err := registry.Decode(dec, name)
	if err == nil {
		if err := blackmagic.AssignIfCompatible(dst, decoded); err != nil {
			return fmt.Errorf(`failed to assign decoded value for %q: %w`, name, err)
		}
		return nil
	}

	return fmt.Errorf(`failed to decode field %q: %w`, name, err)
}
