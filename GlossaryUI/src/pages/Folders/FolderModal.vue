<template>
  <b-modal id="folder-modal" title="New Folder" hide-backdrop>
    <template>
      <label class="font-weight-bold" for="nameInput">Name:* </label>
      <input
        type="text"
        class="form-control"
        v-model="folders.folderName"
        id="nameInput"
      />
      <div class="row modal-body d-flex justify-content-between">
        <div class="col-5">
          <div class="row mt-3">
            <label class="font-weight-bold" for="sourceInput">Source:* </label>
            <select
              type="text"
              class="form-control"
              v-model="folders.source"
              id="sourceInput"
            >
              <option
                v-for="lang in languageOptions"
                :key="lang.id"
                v-bind:value="{ id: lang.id, text: lang.name }"
              >
                {{ itemValue(lang) }}
              </option>
            </select>
          </div>
        </div>
        <div class="col-5">
          <div class="row mt-3">
            <label class="font-weight-bold" for="destinationInput"
              >Destination:*
            </label>
            <select
              type="text"
              class="form-control"
              v-model="folders.destination"
              id="destinationInput"
            >
              <option
                v-for="lang in languageOptions"
                v-bind:value="{ id: lang.id, text: lang.name }"
                :key="lang.id"
              >
                {{ itemValue(lang) }}
              </option>
            </select>
          </div>
        </div>
      </div>
    </template>

    <template #modal-footer>
      <div class="d-flex justify-content-start">
        <div class="btn-margin">
          <b-button class="btn btn-fill float-right  btn-sm  btn-info"  @click.prevent="save"
            >Save</b-button>
          </div>
            <div  class="btn-margin">
            <b-button class="btn btn-fill float-right  btn-danger"  @click.prevent="deleteItem"
                >Delete</b-button>
          </div>
          </div>
     </template>
  </b-modal>
</template>

<script>
import Card from "src/components/Cards/Card.vue";
import axios from "axios";
import URL from "src/util/url";
export default {
  name: "folder-modal",
  components: {
    Card,
  },
  props: ["languageOptions"],
  data() {
    return {
      folders: {
        folderId: 0,
        languageOptions: [],
        folderName: "",
        source: "",
        destination: "",
        BASE_URL: URL.LOCAL_BASE,
      },
    };
  },

  methods: {
    mounted() {
      this.$emit("newFolder", this.newFolder);
    },
    itemValue(folder) {
      return folder["name"];
    },
    newFolder() {
      this.folders["folderId"] = null;
      this.$bvModal.show("folder-modal");
    },
    getAll() {
      axios
        .get(`${this.folders.BASE_URL}/folders-and-languages`)
        .then((response) => {
          this.folders["folderOptions"] = response.data.folders;
          this.folders["languageOptions"] = response.data.languages;
        });
    },
    add() {
      this.$bvModal.show("folder-modal");
    },
    save() {
      let id = this.folders["folderId"];
      if (id) {
        this.edit(id);
      } else {
        this.doSave();
      }
    },
    doSave() {
      let name = this.folders["folderName"];
      let source = this.folders["source"];
      let destination = this.folders["destination"];
      let folder = {
        name: name,
        sourceLanguage: {
          id: source.id,
          name: source.text,
        },
        destinationLanguage: {
          id: destination.id,
          name: destination.text,
        },
      };
      axios.post(`${this.folders.BASE_URL}/folders`, folder).then(() => {
        this.$bvModal.hide("folder-modal");
        this.$emit("getAll");
      });
    },
    showEditModal(id) {
      axios.get(`${this.folders.BASE_URL}/folders/${id}`).then((response) => {
        this.populate(response);

        this.$bvModal.show("folder-modal");

      })

    },
    populate(response) {
      if (response) {
         this.folders["folderId"] = response.data.id;
         this.folders["folderName"] = response.data.name;
         this.folders["source"] = {'id': response.data.sourceLanguage.id, text: response.data.sourceLanguage.name }
         this.folders["destination"] = {'id': response.data.destinationLanguage.id, text: response.data.destinationLanguage.name }
      } else {
         this.folders["folderId"] = "";
         this.folders["folderName"] = "";
         this.folders["source"] = "";
         this.folders["destination"] = "";
      }
    },
    edit(id) {
      let value = this.types["itemName"];
      axios
        .put(`${this.folders.BASE_URL}/types/${id}`, {
          value: value,
          id: id,
        })
        .then(() => {
          this.$bvModal.hide("type-modal");
          this.$emit("getAll");
        });
    },
    showEdit(item) {
      this.types["folderName"] = item.name;
      this.types["folderId"] = item.id;
      this.$bvModal.show("type-modal");
    },
    deleteItem(item) {
      debugger
      axios.delete(`${this.folders.BASE_URL}/folders/${this.folders["folderId"]}`).then(() => {
        this.$emit("getAll");
        this.$bvModal.hide("type-modal");
      });
    },
  },
};
</script>
<style>
.btn-margin {
  margin: 5px;
}
</style>