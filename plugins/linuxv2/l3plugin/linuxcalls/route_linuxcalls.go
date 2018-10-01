//  Copyright (c) 2018 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// +build !windows,!darwin

package linuxcalls

import (
	"time"

	"github.com/vishvananda/netlink"
)

// AddStaticRoute creates the new static route
func (h *NetLinkHandler) AddStaticRoute(route *netlink.Route) error {
	defer func(t time.Time) {
		h.stopwatch.TimeLog("add-static-route").LogTimeEntry(time.Since(t))
	}(time.Now())

	return netlink.RouteAdd(route)
}

// ReplaceStaticRoute removes the static route
func (h *NetLinkHandler) ReplaceStaticRoute(route *netlink.Route) error {
	defer func(t time.Time) {
		h.stopwatch.TimeLog("replace-static-route").LogTimeEntry(time.Since(t))
	}(time.Now())

	return netlink.RouteReplace(route)
}

// DelStaticRoute removes the static route
func (h *NetLinkHandler) DelStaticRoute(route *netlink.Route) error {
	defer func(t time.Time) {
		h.stopwatch.TimeLog("del-static-route").LogTimeEntry(time.Since(t))
	}(time.Now())

	return netlink.RouteDel(route)
}

// GetStaticRoutes reads all configured static routes with the given outgoing
// interface.
// <interfaceIdx> works as filter, if set to zero, all routes in the namespace
// are returned.
func (h *NetLinkHandler) GetStaticRoutes(interfaceIdx int) (v4Routes, v6Routes []netlink.Route, err error) {
	defer func(t time.Time) {
		h.stopwatch.TimeLog("get-static-routes").LogTimeEntry(time.Since(t))
	}(time.Now())

	var link netlink.Link
	if interfaceIdx != 0 {
		// netlink.RouteList reads only link index
		link = &netlink.Dummy{LinkAttrs: netlink.LinkAttrs{Index: interfaceIdx}}
	}

	v4Routes, err = netlink.RouteList(link, netlink.FAMILY_V4)
	if err != nil {
		return
	}
	v6Routes, err = netlink.RouteList(link, netlink.FAMILY_V6)
	return
}