// Set constants
const inputLangaugeSwitcherId = `language`
const inputPageTitleId = `title`
const inputPageCategoryId = `category`
const inputPageSubcategoryId = `subcategory`
const inputPageTextId = `text`

const inputLangaugeSwitcher = document.querySelector(`#${inputLangaugeSwitcherId}`)
const inputPageTitle = document.querySelector(`#${inputPageTitleId}`)
const inputPageCategory = document.querySelector(`#${inputPageCategoryId}`)
const inputPageSubcategory = document.querySelector(`#${inputPageSubcategoryId}`)
const inputPageText = document.querySelector(`#${inputPageTextId}`)

let selectedLanguage = ""
let messageConfirmLanguageChange = "You will lose your changes if you " +
    "change the language without saving.\nDo you want to proceed?"

// Add language change event listener
inputLangaugeSwitcher.addEventListener('change', e => {
    let currentLanguage = e.target.value
    let itemCollection = {
        [inputPageTitleId]: inputPageTitle, 
        [inputPageCategoryId]: inputPageCategory, 
        [inputPageSubcategoryId]: inputPageSubcategory, 
        [inputPageTextId]: inputPageText
    }

    // Check if there had been changes
    // Prompt if the user wants to overwrite changes
    if(selectedLanguage !== "" && selectedLanguage != currentLanguage) {
       if(!window.confirm(messageConfirmLanguageChange)) {
            inputLangaugeSwitcher.value = selectedLanguage
            return
       } 
    }

    // Populate language texts into form fields
    Object.keys(itemCollection).forEach(key => {
        let fieldValue = ""
        document.pageData[key].forEach(item => {
            if(item["lang"] == currentLanguage) {
                fieldValue = item["text"]
            }
        })

        itemCollection[key].value = fieldValue
        if(key == inputPageTextId) {
            tinymce.get(inputPageTextId).setContent(fieldValue)
        }
    })

    selectedLanguage = currentLanguage
})

