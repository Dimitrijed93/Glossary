<template>
  <table class="table">
    <thead>
      <slot name="columns">
        <tr>
          <th v-for="column in columns" :key="column">{{column}}</th>
          <th></th>
        </tr>
      </slot>
    </thead>
    <tbody>
    <tr v-for="(item, index) in data" :key="index">
      <slot :row="item">
        <td v-for="column in columns" :key="column" v-if="hasValue(item, column)">{{itemValue(item, column)}}</td>
        <td class="removeIcon"><button class="btn btn-info pull-right" v-on:click="$emit('editItem', item)">Edit</button></td>
        <td class="removeIcon"><button class="btn btn-danger pull-right" v-on:click="$emit('deleteItem', item)">Delete</button></td>

      </slot>
    </tr>
    </tbody>
  </table>
</template>
<script>
export default {
  name: "l-table",
  props: {
    columns: Array,
    data: Array,
  },
  methods: {
    hasValue(item, column) {
      return item[column.toLowerCase()] !== "undefined";
    },
    itemValue(item, column) {
      return item[column.toLowerCase()];
    },
  },
};
</script>
<style>
</style>
