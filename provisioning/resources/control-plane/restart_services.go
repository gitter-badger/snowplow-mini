/**
 * Copyright (c) 2016-2017 Snowplow Analytics Ltd.
 * All rights reserved.
 *
 * This program is licensed to you under the Apache License Version 2.0,
 * and you may not use this file except in compliance with the Apache
 * License Version 2.0.
 * You may obtain a copy of the Apache License Version 2.0 at
 * http://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the Apache License Version 2.0 is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied.
 *
 * See the Apache License Version 2.0 for the specific language
 * governing permissions and limitations there under.
 */

package main

import (
	"errors"
	"os/exec"
)

func restartService(service string) error {
	var initMap = map[string]string{
		"streamCollector": config.Inits.StreamCollector,
		"streamEnrich":    config.Inits.StreamEnrich,
		"esLoaderGood":    config.Inits.EsLoaderGood,
		"esLoaderBad":     config.Inits.EsLoaderBad,
		"iglu":            config.Inits.Iglu,
		"caddy":           config.Inits.Caddy,
	}

	if val, ok := initMap[service]; ok {
		restartCommand := []string{"service", val, "restart"}

		cmd := exec.Command("/bin/bash", restartCommand...)
		err := cmd.Run()
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("unrecognized service")
}

func restartSPServices() error {
	err := restartService("streamCollector")
	if err != nil {
		return err
	}

	err = restartService("streamEnrich")
	if err != nil {
		return err
	}

	err = restartService("esLoaderGood")
	if err != nil {
		return err
	}

	err = restartService("esLoaderBad")
	if err != nil {
		return err
	}

	return nil
}
