<template>
    <v-card class="mx-auto">
        <v-card-title>
            {{ Filename }}
        </v-card-title>
        <v-card-subtitle>
            <v-row>
                <v-col cols="2">
                    <span class="subheading">ID: {{ ID }}</span>

                </v-col>
                <v-col>
                    <span class="subheading"> Size: {{ Size }}</span>
                </v-col>
            </v-row>
        </v-card-subtitle>

        <v-card-text>
            {{ Description }}
        </v-card-text>

        <v-divider class="mx-4"></v-divider>

        <v-card-text>
            <div class="text-h6">
                Benchmarks:
            </div>
            <v-chip-group column multiple v-model="selectionTest" v-show="test">
                <v-chip v-for="benchmark in Benchmarks" :key="benchmark + 'test'" :value="benchmark" color="green"
                    :disabled="isDisabled(benchmark, true)">
                    {{ benchmark }}
                </v-chip>
            </v-chip-group>
            <v-chip-group column multiple v-model="selectionRef" v-show="!test">
                <v-chip v-for="benchmark in Benchmarks" :key="benchmark" :value="benchmark" color="green"
                    :disabled="isDisabled(benchmark, false)">
                    {{ benchmark }}
                </v-chip>
            </v-chip-group>
        </v-card-text>
        <v-card-actions>
            <div class="text-h6">
                Size: {{ test ? "test" : "ref" }}
            </div>
            <v-spacer></v-spacer>
            <v-btn icon="mdi-swap-horizontal" @click="test = !test"></v-btn>

        </v-card-actions>


        <v-divider class="mx-4"></v-divider>



        <v-card-actions>
            <v-btn @click="newTask" color="blue-accent-2" variant="outlined">
                New Task
            </v-btn>
            <v-btn @click="download" color="blue-accent-2" variant="outlined">
                Download Data
            </v-btn>
            <v-btn @click="openconfig" color="blue-accent-2" variant="outlined">
                View Config
            </v-btn>
            <v-btn @click="intrate" color="deep-purple-accent-4">
                intrate
            </v-btn>
            <v-btn @click="intspeed" color="red-darken-1">
                intspeed
            </v-btn>
            <v-btn @click="fprate" color="deep-purple-accent-4">
                fprate
            </v-btn>
            <v-btn @click="fpspeed" color="red-darken-1">
                fpspeed
            </v-btn>
            <v-btn @click="clearselect" color="blue">
                clear
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn icon="mdi-refresh" @click="emit('refresh')"></v-btn>

            <v-btn :icon="show ? 'mdi-chevron-up' : 'mdi-chevron-down'" @click="show = !show; test = !!test"></v-btn>
        </v-card-actions>

        <v-expand-transition>
            <div v-show="show">
                <v-divider></v-divider>
                <v-col v-for="qemutask in QemuTasks" :key="qemutask.ID" cols="12"
                    v-show="test ? (qemutask.Type == 'test' && (selectionTest.includes(qemutask.Benchmark) || selectionTest.length == 0)) : (qemutask.Type == 'ref' && (selectionRef.includes(qemutask.Benchmark) || selectionRef.length == 0))">
                    <qemutask-card v-bind="qemutask" />
                </v-col>
            </div>
        </v-expand-transition>

    </v-card>

</template>
<script setup lang="ts">


const props = defineProps<{
    ID: number
    Filename: string
    Benchmarks: [string]
    Size: number
    Description: string
    QemuTasks: [any]
    // n: number
}>()

const emit = defineEmits(['newtask', 'refresh'])

const intrateBenchmarks = [
    "500.perlbench_r",
    "502.gcc_r",
    "505.mcf_r",
    "520.omnetpp_r",
    "523.xalancbmk_r",
    "525.x264_r",
    "531.deepsjeng_r",
    "541.leela_r",
    "548.exchange2_r",
    "557.xz_r",
]
const intspeedBenchmarks = [
    "600.perlbench_s",
    "602.gcc_s",
    "605.mcf_s",
    "620.omnetpp_s",
    "623.xalancbmk_s",
    "625.x264_s",
    "631.deepsjeng_s",
    "641.leela_s",
    "648.exchange2_s",
    "657.xz_s"
]
const fprateBenchmarks = [
    "503.bwaves_r",
    "507.cactuBSSN_r",
    "508.namd_r",
    "510.parest_r",
    "511.povray_r",
    "519.lbm_r",
    "521.wrf_r",
    "526.blender_r",
    "527.cam4_r",
    "538.imagick_r",
    "544.nab_r",
    "549.fotonik3d_r",
    "554.roms_r",
]
const fpspeedBenchmarks = [
    "603.bwaves_s",
    "607.cactuBSSN_s",
    "619.lbm_s",
    "621.wrf_s",
    "627.cam4_s",
    "628.pop2_s",
    "638.imagick_s",
    "644.nab_s",
    "649.fotonik3d_s",
    "654.roms_s",
]

const show = ref(false)
const test = ref(false)
const selectionTest = ref([] as any[])
const selectionRef = ref([] as any[])

function intrate() {
    for (let b of intrateBenchmarks) {
        if (isDisabled(b, test.value)) {
            continue
        }
        if (test.value) {
            if (!(b in selectionTest.value)) {
                selectionTest.value.push(b)
            }
        } else {
            if (!(b in selectionRef.value)) {
                selectionRef.value.push(b)
            }
        }
    }
}
function intspeed() {
    for (let b of intspeedBenchmarks) {
        if (isDisabled(b, test.value)) {
            continue
        }
        if (test.value) {
            if (!(b in selectionTest.value)) {
                selectionTest.value.push(b)
            }
        } else {
            if (!(b in selectionRef.value)) {
                selectionRef.value.push(b)
            }
        }
    }
}
function fprate() {
    for (let b of fprateBenchmarks) {
        if (isDisabled(b, test.value)) {
            continue
        }
        if (test.value) {
            if (!(b in selectionTest.value)) {
                selectionTest.value.push(b)
            }
        } else {
            if (!(b in selectionRef.value)) {
                selectionRef.value.push(b)
            }
        }
    }
}
function fpspeed() {
    for (let b of fpspeedBenchmarks) {
        if (isDisabled(b, test.value)) {
            continue
        }
        if (test.value) {
            if (!(b in selectionTest.value)) {
                selectionTest.value.push(b)
            }
        } else {
            if (!(b in selectionRef.value)) {
                selectionRef.value.push(b)
            }
        }
    }
}
function clearselect() {
    if (test.value) {
        selectionTest.value = []
    } else {
        selectionRef.value = []
    }
}



function isDisabled(benchmark: string, testGroup: boolean) {
    return false
    const benchmarkType = testGroup ? 'test' : 'ref'
    for (let qemutask of props.QemuTasks) {
        if (qemutask.Type == benchmarkType && qemutask.Benchmark == benchmark) {
            return true
        }
    }
    return false
}

function download() {
    let selectBenchmark = test.value ? selectionTest.value : selectionRef.value
    if (selectBenchmark.length == 0) {
        selectBenchmark = props.Benchmarks
    }
    const header = "benchmark,instructions\n"
    let data = []
    for (let b of selectBenchmark) {
        let find = false
        for (let qemutask of props.QemuTasks) {
            if (qemutask.Benchmark == b) {
                find = true
                if (qemutask.Status == 'success') {
                    data.push(`${qemutask.Name},${qemutask.Inscount}\n`)
                } else {
                    data.push(`${qemutask.Name},${qemutask.Status}\n`)
                }
            }
        }
        if (!find) {
            data.push(`${b}, not run\n`)
        }
    }
    data.sort()
    // const blob = new Blob([header, ...data], { type: 'text/csv' })
    // const url = URL.createObjectURL(blob)
    var element = document.createElement('a');
    element.setAttribute('href', 'data:text/csv;charset=utf-8,' + encodeURIComponent(header + data.join('')));
    element.setAttribute('download', props.Filename + '_' + test ? 'test' : 'ref' + '.csv');

    element.style.display = 'none';
    document.body.appendChild(element);

    element.click();

    document.body.removeChild(element);


}

function openconfig() {
    window.open("/api/v1/tarball/config/" + props.ID)
}

function newTask() {
    console.log('new task')
    console.log(props.ID)
    console.log(test.value)

    if (test.value) {
        console.log(selectionTest.value)
    } else {
        console.log(selectionRef.value)
    }
    fetch('/api/v1/task', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            tarball_id: props.ID,
            type: test.value ? 'test' : 'ref',
            benchmarks: test.value ? selectionTest.value : selectionRef.value
        })
    })
        .then(res => res.json())
        .then(data => {
            console.log(data)
            emit('newtask')
        })
}
</script>