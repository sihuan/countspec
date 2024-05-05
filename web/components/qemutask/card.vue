<template>
    <v-card :color="color" variant="tonal" class="mx-auto">
        <v-card-item>
            <div>
                <v-row>
                    <v-col cols="3">
                        <div class="text-h6">
                            {{ Name }}
                        </div>
                    </v-col>
                    <v-col cols="3" v-show="Status === 'success'">
                        <div class="text-subtitle-1">
                            instructions:  {{ Inscount }}
                        </div>
                    </v-col>
                    <v-col cols="3" v-show="Status != 'success'">
                        <div class="text-subtitle-1">
                            Status:  {{Status }}
                        </div>
                    </v-col>
                    <v-col @click="stdout" cols="3">
                        <div class="text-subtitle-1">
                            stdout:  {{ Stdout }}
                        </div>
                    </v-col>
                    <v-col  @click="stderr" cols="3">
                        <div class="text-subtitle-1">
                            stderr:  {{ Stderr }}
                        </div>
                    </v-col>

                </v-row>


                <!-- <div class="text-caption">Greyhound divisely hello coldly fonwderfully</div> -->
            </div>
        </v-card-item>
    </v-card>

</template>
<script setup lang="ts">
const props = defineProps<{
    ID: number
    Name: string
    BenchmarkName: string
    Status: string
    Inscount: number
    Stderr: string
    Stdout: string
}>()

function stdout() {
    window.open('/api/v1/task/stdout/' + props.ID)
}
function stderr() {
    window.open('/api/v1/task/stderr/' + props.ID)
}

const color = computed(() => {
    if (props.Status === 'running') {
        return 'blue'
    } else if (props.Status === 'success') {
        return 'green'
    } else if (props.Status === 'failed') {
        return 'red'
    } else {
        return 'grey'
    }
})
</script>