<template>
    <div class="grid grid-cols-1 grid-rows-2">
        <div>
            <div v-for="chan of chans" class="flex flex-row">
                <div class="mx-2">{{ chan.name }}</div>
                <div class="mx-2">{{ chan.id }}</div>
                <div class="mx-2">{{ chan.userount }}</div>
            </div>
        </div>
        <div>
            <button @click="fetchChannels" class="bg-red-900">Refresh</button>
        </div>
    </div>
</template>


<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { Channel } from './channel';
import axios from 'axios';
import config from "../../config.ts"


let chans = ref<Array<Channel>>([])


async function fetchChannels(): Promise<void> {
    let resp = await axios.get(`${config.BASE_URL}/channel`)
    for(let ch of resp.data.data) {
        let chan = new Channel(ch.channel_id, ch.user_count, ch.channel_name)
        chans.value?.push(chan)
    }
}

onMounted( () => {
    const chs: Channel[] = [
        new Channel(1, 5, "reaper 1"),
        new Channel(2, 4, "reaper 2"),
        new Channel(3, 7, "reaper 3"),
    ]
    // await fetchChannels()
    chans.value?.push(...chs)

    console.log(chans.value?.length)
})

</script>