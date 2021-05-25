<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row"></div>
      <div class="row">
        <div class="col-12">
          <card class="card-plain">
            <template slot="header">
              <h4 class="card-title">Types</h4>
            </template>
            <div class="table-responsive">
              <l-table
                class="table-hover"
                :columns="types.columns"
                :data="types.data"
                v-on:editItem="showEdit"
                v-on:deleteItem="deleteItem"
              >
              </l-table>
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
          </card>
        </div>
      </div>
    </div>

    <b-modal id="type-modal" hide-footer>
      <template #modal-title> <b>New Type</b></template>
      <div class="row mt-3">
        <label for="nameInput">Name: * </label>
        <input
          type="text"
          class="form-control"
          v-model="types.itemName"
          id="nameInput"
        />
      </div>
      <div class="row mt-3">
        <b-button class="btn btn-primary" block @click.prevent="save"
          >Save</b-button
        >
      </div>
    </b-modal>
  </div>
</template>

<script>
import LTable from "src/components/Table.vue";
import Card from "src/components/Cards/Card.vue";
import axios from "axios";
const tableColumns = ["Value"];
export default {
  components: {
    LTable,
    Card,
  },
  data() {
    return {
      types: {
        columns: [...tableColumns],
        data: [],
        itemName: "",
        itemId: "",
      },
    };
  },
  mounted() {
    this.getAll();
  },
  methods: {
    getAll() {
      axios.get("http://localhost:9090/types").then((response) => {
        this.types["data"] = response.data;
      });
    },
    add() {
      this.$bvModal.show("type-modal");
    },
    save() {
      let id = this.types["itemId"];
      if (id) {
        this.edit(id);
      } else {
        this.doSave();
      }
    },
    doSave() {
      let value = this.types["itemName"];
      axios
        .post("http://localhost:9090/types", {
          value,
        })
        .then(() => {
          this.$bvModal.hide("type-modal");
          this.getAll();
        });
    },
    edit(id) {
      let value = this.types["itemName"];
      axios
        .put(`http://localhost:9090/types/${id}`, {
          value: value,
          id: id,
        })
        .then(() => {
          this.$bvModal.hide("type-modal");
          this.getAll();
        });
    },
    showEdit(item) {
      this.types["itemName"] = item.value;
      this.types["itemId"] = item.id;
      this.$bvModal.show("type-modal");
    },
    deleteItem(item) {
      axios.delete(`http://localhost:9090/types/${item.id}`).then(() => {
        this.getAll();
      });
    },
  },
};
</script>
<style>
</style>
