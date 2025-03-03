/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import (
	"github.com/liulingba/chaosblade-exec-os/exec/cpu"
	"github.com/liulingba/chaosblade-exec-os/exec/disk"
	"github.com/liulingba/chaosblade-exec-os/exec/file"
	"github.com/liulingba/chaosblade-exec-os/exec/kernel"
	"github.com/liulingba/chaosblade-exec-os/exec/mem"
	"github.com/liulingba/chaosblade-exec-os/exec/network"
	"github.com/liulingba/chaosblade-exec-os/exec/process"
	"github.com/liulingba/chaosblade-exec-os/exec/script"
	"github.com/liulingba/chaosblade-exec-os/exec/systemd"
	"github.com/liulingba/chaosblade-exec-os/exec/time"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

// GetAllExpModels returns the experiment model specs in the project.
// Support for other project about chaosblade
func GetAllExpModels() []spec.ExpModelCommandSpec {
	return []spec.ExpModelCommandSpec{
		cpu.NewCpuCommandModelSpec(),
		mem.NewMemCommandModelSpec(),
		process.NewProcessCommandModelSpec(),
		network.NewNetworkCommandSpec(),
		disk.NewDiskCommandSpec(),
		script.NewScriptCommandModelSpec(),
		file.NewFileCommandSpec(),
		kernel.NewKernelInjectCommandSpec(),
		systemd.NewSystemdCommandModelSpec(),
		time.NewTimeCommandSpec(),
	}
}
