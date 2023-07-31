<script>
    import { Input } from "sveltestrap";
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = "";
	export let table_body_font = "";
	export let token = "";
	export let listHome = [];
	export let listCurr = [];
	export let totalrecord = 0;
    let dispatch = createEventDispatcher();
	let title_page = "MASTER"
    let sData = "";
    let myModal_newentry = "";
    let flag_id_field = false;
    let flag_btnsave = true;
    let idcurr_field = "";
    let name_field = "";
    let owner_field = "";
    let email_field = "";
    let phone1_field = "";
    let phone2_field = "";
    let note_field = "";
    let status_field = "";
    let create_field = "";
    let update_field = "";
    let idrecord = "";
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    
    const NewData = (e,id,idcurr,name,owner,email,phone1,phone2,note,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            flag_id_field = true;
            idrecord = id
            idcurr_field = idcurr;
            name_field = name;
            owner_field = owner;
            email_field = email;
            phone1_field = phone1;
            phone2_field = phone2;
            note_field = note;
            status_field = status;
            create_field = create;
            update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const GenerateString = (e) => {
        idrecord = genRandomString(e)
    };
    
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(idrecord.length < 4){
                flag = false
                msg += "The ID is maxlength 4\n"
            }
            if(idcurr_field == ""){
                flag = false
                msg += "The Default Curreny is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(phone1_field == ""){
                flag = false
                msg += "The Phone 1 is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(idcurr_field == ""){
                flag = false
                msg += "The Default Curreny is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(phone1_field == ""){
                flag = false
                msg += "The Phone 1 is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/mastersave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"MASTER-SAVE",
                    master_id: idrecord,
                    master_idcurr: idcurr_field,
                    master_name: name_field,
                    master_owner: owner_field,
                    master_phone1: phone1_field,
                    master_phone2: phone2_field,
                    master_email: email_field,
                    master_note: note_field,
                    master_status: status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    
    function clearField(){
        idrecord = "";
        idcurr_field = "";
        name_field = "";
        owner_field = "";
        email_field = "";
        phone1_field = "";
        phone2_field = "";
        note_field = "";
        status_field = "";
        flag_id_field = false
        create_field = "";
        update_field = "";
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","","","","","","","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterTafsirMimpi = [];
                listHome = [];
                const tafsir = {
                    searchTafsirMimpi,
                };
                dispatch("handleTafsirMimpi", tafsir);
        }  
    };
    function status(e){
        let result = "DEACTIVE"
        if(e == "Y"){
            result = "ACTIVE"
        }
        return result
    }
    function uperCase(element) {
        function onInput(event) {
            element.value = element.value.toUpperCase();
        }
        element.addEventListener("input", onInput);
        return {
            destroy() {
                element.removeEventListener("input", onInput);
            },
        };
    }
    function genRandomString(length) {
        var chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
        var charLength = chars.length;
        var result = '';
        for ( var i = 0; i < length; i++ ) {
            result += chars.charAt(Math.floor(Math.random() * charLength));
        }
        return result;
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            
            <Button
                on:click={callFunction}
                button_function="NEW"
                button_title="New"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchHome}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Master"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">START</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">END</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CODE</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CURR</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">MASTER</th>
                                <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">OWNER</th>
                                <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                                <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">EMAIL</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                //e,id,idcurr,name,owner,email,phone1,phone2,note,status,create,update
                                                NewData("Edit",rec.home_id, rec.home_idcurr,
                                                rec.home_name,rec.home_owner,rec.home_email,rec.home_phone1,rec.home_phone2,rec.home_note,rec.home_status,
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td nowrap style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_start}</td>
                                    <td nowrap style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_end}</td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_idcurr}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_name}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_owner}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        {rec.home_phone1} / {rec.home_phone2}
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_email}</td>
                                </tr>
                            {/each}
                        </tbody>
                        {:else}
                        <tbody>
                            <tr>
                                <td colspan="20">
                                    <center>
                                        <Loader />
                                    </center>
                                </td>
                            </tr>
                        </tbody>
                        {/if} 
                    </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="modalentrycrud"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">CODE</label>
                    <div class="input-group mb-3">
                        <input bind:value={idrecord}
                            use:uperCase
                            class="required form-control"
                            maxlength="3"
                            disabled
                            type="text"
                            placeholder="CODE"/>
                        {#if sData != "Edit"}
                        <button on:click={() => {
                                GenerateString(4);
                            }}
                            type="button" class="btn btn-info">Generate</button>
                        {/if}
                      </div>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Default Currency</label>
                    <select
                        bind:value="{idcurr_field}" 
                        name="departement" id="departement" 
                        class="required form-control">
                        {#each listCurr as rec}
                        <option value="{rec.curr_id}">{rec.curr_id}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input bind:value={name_field}
                        class="required"
                        type="text"
                        placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Owner</label>
                    <Input bind:value={owner_field}
                        class="required"
                        type="text"
                        placeholder="Owner"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Email</label>
                    <Input bind:value={email_field}
                        class=""
                        type="text"
                        placeholder="Email"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 1</label>
                    <Input bind:value={phone1_field}
                        class="required"
                        type="text"
                        placeholder="Phone 1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 2</label>
                    <Input bind:value={phone2_field}
                        class=""
                        type="text"
                        placeholder="Phone 2"/>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Note</label>
                    <textarea style="height: 100px;resize: none;" bind:value={note_field} class="form-control"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={status_field}>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                    <div class="mb-3">
                        <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                            Create : {create_field}<br />
                            Update : {update_field}
                        </div>
                    </div>
                {/if}
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>



