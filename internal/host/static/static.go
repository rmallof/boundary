package static

import (
	"errors"

	"github.com/hashicorp/watchtower/internal/db"
	"github.com/hashicorp/watchtower/internal/host/static/store"
)

// A HostCatalog contains static hosts and static host sets.
type HostCatalog struct {
	*store.HostCatalog
	tableName string `gorm:"-"`
}

// NewHostCatalog creates a new in memory HostCatalog assigned to scopeId.
// Name and description are the only valid options. All other options are
// ignored.
func NewHostCatalog(scopeId string, opt ...Option) (*HostCatalog, error) {
	if scopeId == "" {
		return nil, errors.New("empty scopeId")
	}
	id, err := db.NewPublicId("sthc")
	if err != nil {
		return nil, err
	}
	opts := getOpts(opt...)
	hc := &HostCatalog{
		HostCatalog: &store.HostCatalog{
			ScopeId:     scopeId,
			Name:        opts.withName,
			Description: opts.withDescription,
			PublicId:    id,
		},
	}
	return hc, nil
}

// A Host contains a static address.
type Host struct {
	*store.Host
	tableName string `gorm:"-"`
}

// NewHost creates a new in memory Host for address assigned to catalogId.
// Name and description are the only valid options. All other options are
// ignored.
func NewHost(catalogId, address string, opt ...Option) (*Host, error) {
	if catalogId == "" {
		return nil, errors.New("empty catalogId")
	}
	if address == "" {
		return nil, errors.New("empty address")
	}
	id, err := db.NewPublicId("sth")
	if err != nil {
		return nil, err
	}
	opts := getOpts(opt...)
	host := &Host{
		Host: &store.Host{
			StaticHostCatalogId: catalogId,
			Address:             address,
			Name:                opts.withName,
			Description:         opts.withDescription,
			PublicId:            id,
		},
	}
	return host, nil
}

// A HostSet contains a static address.
type HostSet struct {
	*store.HostSet
	tableName string `gorm:"-"`
}

// NewHostSet creates a new in memory HostSet assigned to catalogId.
// Name and description are the only valid options. All other options are
// ignored.
func NewHostSet(catalogId string, opt ...Option) (*HostSet, error) {
	if catalogId == "" {
		return nil, errors.New("empty catalogId")
	}
	id, err := db.NewPublicId("sths")
	if err != nil {
		return nil, err
	}
	opts := getOpts(opt...)
	set := &HostSet{
		HostSet: &store.HostSet{
			StaticHostCatalogId: catalogId,
			Name:                opts.withName,
			Description:         opts.withDescription,
			PublicId:            id,
		},
	}
	return set, nil
}

// A HostSetMember contains a static address.
type HostSetMember struct {
	*store.HostSetMember
	tableName string `gorm:"-"`
}

// NewHostSetMember creates a new in memory HostSetMember assigned to catalogId.
// Name and description are the only valid options. All other options are
// ignored.
func NewHostSetMember(hostSetId, hostId string, opt ...Option) (*HostSetMember, error) {
	if hostSetId == "" {
		return nil, errors.New("empty hostSetId")
	}
	if hostId == "" {
		return nil, errors.New("empty hostId")
	}
	member := &HostSetMember{
		HostSetMember: &store.HostSetMember{
			StaticHostSetId: hostSetId,
			StaticHostId:    hostId,
		},
	}
	return member, nil
}