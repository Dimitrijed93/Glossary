<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row"></div>
      <div class="row">
        <div class="col-12">
          <card class="card-plain">
            <template slot="header">
              <h4 class="card-title">Languages</h4>
            </template>
            <div class="table-responsive">
              <l-table
                class="table-hover"
                :columns="languages.columns"
                :data="languages.data"
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

    <b-modal id="language-modal" hide-footer>
      <template #modal-title> <b>New Language</b></template>
      <div class="row mt-3">
        <label for="nameInput">Name: * </label>
        <input
          type="text"
          class="form-control"
          v-model="languages.itemName"
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
const tableColumns = ["Name"];
export default {
  components: {
    LTable,
    Card,
  },
  data() {
    return {
      languages: {
        columns: [...tableColumns],
        data: [],
        itemName: "",
        itemId: "",
        BASE_URL: URL.LOCAL_BASE

      },
    };
  },
  mounted() {
    this.getAll();
  },
  methods: {
    getAll() {
      axios.get(`${this.languages.BASE_URL}/languages`).then((response) => {
        this.languages["data"] = response.data;
      });
    },
    add() {
      this.$bvModal.show("language-modal");
    },
    save() {
      let id = this.languages["itemId"];
      if (id) {
        this.edit(id);
      } else {
        this.doSave();
      }
    },
    doSave() {
      let name = this.languages["itemName"];
      axios
        .post(`${this.languages.BASE_URL}/languages`, {
          name,
        })
        .then(() => {
          this.$bvModal.hide("language-modal");
          this.getAll();
        });
    },
    edit(id) {
      let name = this.languages["itemName"];
      axios
        .put(`${this.languages.BASE_URL}/languages/${id}`, {
          name: name,
          id: id,
        })
        .then(() => {
          this.$bvModal.hide("language-modal");
          this.getAll();
        });
    },
    showEdit(item) {
      this.languages["itemName"] = item.name;
      this.languages["itemId"] = item.id;
      this.$bvModal.show("language-modal");
    },
    deleteItem(item) {
      debugger
      axios.delete(`${this.languages.BASE_URL}/languages/${item.id}`).then(() => {
        this.getAll();
      });
    },
  },
};
</script>
<style>
</style>
