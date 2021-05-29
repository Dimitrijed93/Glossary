<template>
    <card class="strpied-tabled-with-hover" >
      <h4>{{data}}</h4>
    </card>
</template>

<script>
import Card from "src/components/Cards/Card.vue";
import axios from "axios";
import URL from "src/util/url";
export default {
  name: "entry",
  components: {
    Card,
  },
  props: ["data"],
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
      debugger
      // this.$emit("newFolder", this.newFolder);
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