<template>
    <div class="grid grid-cols-1 grid-rows-2 text-base text-text max-w-fit">
        <div class="flex items-center justify-center font-bold">
            Add channel
        </div>
        <div class="grid grid-cols-1 my-2">
            <label>Link</label>
            <input @input="checkInput" type="text" class="text-dark-main" placeholder="https://t.me/Reaperhome">
            <p v-if="hasError"></p>
        </div>
        <div>
            <button @click="addChannel" class="btn bg-dark-main text-base">Add</button>
        </div>
    </div>
</template>


<script setup lang="ts">
import { onMounted } from 'vue';
import axios from 'axios';
import config from "../../config.ts"

let hasError = false

async function addChannel(): Promise<void> {
    let resp = await axios.post(`${config.BASE_URL}/channel`)
}

function checkInput(event) {
    let val: string = event.target.value
    console.log(val)
}

onMounted( () => {
})

</script>