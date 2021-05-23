<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row">
        <div class="col-12">
          <div class="text-center">
            <button
              type="submit"
              class="btn btn-info btn-fill float-left"
              @click.prevent="add"
            >
              Add
            </button>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-12">
          <card class="card-plain">
            <template slot="header">
              <h4 class="card-title">Types</h4>
              <p class="card-category">Types for Extras</p>
            </template>
            <div class="table-responsive">
              <l-table
                class="table-hover"
                :columns="types.columns"
                :data="types.data"
              >
              </l-table>
            </div>
          </card>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import LTable from "src/components/Table.vue";
import Type from "src/models/Type.js";
import Card from "src/components/Cards/Card.vue";
import axios from "axios";
const tableColumns = ["Value"];
const tableData = [new Type(1, "test")];
export default {
  components: {
    LTable,
    Card,
  },
  data() {
    return {
      types: {
        columns: [...tableColumns],
        data: [...tableData],
      },
    };
  },
  mounted() {
  axios.get("http://localhost:9090/types")
    .then(response => {
      this.types = response.data;
      console.log(JSON.stringify(response.data))
    })
  },
  methods: {
    add() {
      alert("Your data: " + JSON.stringify(this.user));
    },
  },
};
</script>
<style>
</style>
