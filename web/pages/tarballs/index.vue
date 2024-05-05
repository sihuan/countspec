<template>
    <v-container>
        <v-row>
            <v-col v-for="tarball in tarballs" :key="tarball.ID" cols="12">
                <tarball-card v-bind="tarball" @newtask="refresh" @refresh="refresh"/>
            </v-col>
        </v-row>
    </v-container>
    <!-- <v-fab @click="uploadTarball" color="primary" icon="mdi-plus" app ></v-fab>
     -->
    <v-dialog max-width="500">

        <template v-slot:activator="{ props: activatorProps }">
            <!-- <v-btn v-bind="activatorProps" color="surface-variant" text="Open Dialog" variant="flat"></v-btn>ss -->
            <v-fab v-bind="activatorProps" color="primary" icon="mdi-plus" app></v-fab>

        </template>

        <template v-slot:default="{ isActive }">
            <v-card title="Upload new tarball">
                <v-card-text>
                    <v-file-input v-model="selectedFile" label="Select File" :disabled="uploading"></v-file-input>
                    <v-text-field v-model="description" label="Desc" :disabled="uploading" multi-line
                        rows="3"></v-text-field>
                </v-card-text>
                <v-card-actions v-if="!uploading">
                    <v-spacer></v-spacer>
                    <v-btn color="primary" @click="uploadFile">Upload</v-btn>
                    <v-btn @click="isActive.value = false">Cancel</v-btn>
                </v-card-actions>
                <v-card-actions v-else>
                    <v-progress-linear :indeterminate="uploading" color="primary"></v-progress-linear>
                </v-card-actions>
            </v-card>
        </template>
    </v-dialog>
</template>


<script setup lang="ts">
// export default {
//     data() {
//         return {
//             tarballs: [],
//             selectedFile: null,
//             description: '',
//             uploading: false
//         }
//     },
//     methods: {
//         uploadFile() {
//             if (this.selectedFile) {
//                 var formData = new FormData();
//                 console.log(this.selectedFile)
//                 formData.append('file', this.selectedFile[0]);
//                 // formData.append('description', this.description);
//                 this.uploading = true;
//                 fetch('http://127.0.0.1:9090/v1/tarball', {
//                     method: 'POST',
//                     body: formData,

//                 })
//                     .then(res => res.json())
//                     .then(data => {
//                         if (data.error) {
//                             alert(data.error)
//                         } else { alert('success') }
//                         this.uploading = false;
//                         this.selectedFile = null;
//                         this.description = '';
//                         this.refresh()
//                     })
//             }
//         },
//         refresh() {
//             fetch('/api/v1/tarball')
//                 .then(res => res.json())
//                 .then(data => {
//                     this.tarballs = data.data
//                 })
//         }
//     },
//     mounted() {
//         this.refresh()
//     }
// }


// 从后端获取数据
interface Tarball {
    ID: number
    Filename: string
    Size: number
    Description: string
    ConfigPath: string
    Benchmarks: [string]
    QemuTasks: [any]

}

const tarballs = ref([] as Tarball[])
const selectedFile = ref(null as File[] | null);
const description = ref('');
const uploading = ref(false);

function uploadFile() {
    if (selectedFile.value) {
        var formData = new FormData();
        formData.append('file', selectedFile.value[0]);
        formData.append('description', description.value);
        uploading.value = true;
        fetch('/api/v1/tarball', {
            method: 'POST',
            body: formData,
        })
            .then(res => res.json())
            .then(data => {
                if (data.error) {
                    alert(data.error)
                } else { alert('success') }
                uploading.value = false;
                selectedFile.value = null;
                description.value = '';
                refresh()
            })
    }

};

function sortQemutasks() {
    for (let t of tarballs.value) {
        t.QemuTasks.sort((a, b) => {
            // Status: running, pending, failed, success
            // Name asc
            if (a.Status === b.Status) {
                return a.Name.localeCompare(b.Name)
            } else {
                if (a.Status === 'running') {
                    return -1
                } else if (b.Status === 'running') {
                    return 1
                } else if (a.Status === 'pending') {
                    return -1
                } else if (b.Status === 'pending') {
                    return 1
                } else if (a.Status === 'failed') {
                    return 1
                } else if (b.Status === 'failed') {
                    return -1
                } else if (a.Status === 'success') {
                    return -1
                } else if (b.Status === 'success') {
                    return 1
                }
            }
        })
    }
}


const refresh = () => {
    console.log('refresh')
    fetch('/api/v1/tarball')
        .then(res => res.json())
        .then(data => {
            tarballs.value = data.data
            sortQemutasks()
        })
}

onMounted(async () => {
    refresh()
})

</script>
