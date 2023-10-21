<script lang="ts" setup>
import { reactive } from 'vue';
import { PushCreature } from '../wailsjs/go/main/App';
import { FormKit } from '@formkit/vue'

type Creature = { name: string, image?: string }

let creatures = reactive<Creature[]>([])

const pushNewCreatureToList = async (fields: any) => {
  console.log(fields)
  const fileReader = new FileReader();
  fileReader.readAsDataURL(fields.image[0].file);
  fileReader.onload = () => {
    const payload: Creature = {
      name: fields.name,
      // TODO: handle error
      // @ts-ignore
      image: fileReader.result
    }
    console.log(payload)
    creatures.push(payload)
    PushCreature([{ name: payload.name, image: payload.image?.split(',')[1] }])
  }

}
</script>

<template>
  <div class="container p-4 grid grid-cols-2">
    <div class="grid grid-cols-2 gap-2">
      <div class="w-25 flex justify-around flex-col rounded bg-slate-200 " v-for="creature in creatures">
        <img class="p-3 rounded" v-bind:src="creature.image" />
        <div class="w-full text-center">{{ creature.name }}</div>
      </div>

    </div>

    <FormKit class="pr-2 w-25 flex flex-col justify-around" type="form" @submit="pushNewCreatureToList" :submit-attrs="{
      wrapperClass: 'flex justify-around',
    }">
      <FormKit type="file" name="image" label="Image" accept=".png,.gif,.jpeg" />
      <FormKit validation="required" type="text" name="name" />
    </FormKit>

  </div>
</template>
