import {useState} from 'react';

// Set initial value of form fields
const getInitialFieldsValue = (
    formFields: Partial<Record<FormFields, ModalFormFieldConfig>>,
): Partial<Record<FormFields, string>> => {
    let fields = {};
    Object.keys(formFields).forEach((field) => {
        fields = {
            ...fields,
            [field as FormFields]:
                formFields[field as FormFields]?.value ||
                (field as FormFields === 'timestamp' ? Date.now().toString() : ''),
        };
    });

    return fields as unknown as Partial<Record<FormFields, string>>;
};

/**
 * Filter out all the fields for which validations check required
 * and set empty string as default error message
 */
const getFieldslWhereErrorCheckRequired = (
    formFields: Partial<Record<FormFields, ModalFormFieldConfig>>,
): Partial<Record<FormFields, string>> => {
    let fields = {};
    Object.keys(formFields).forEach((field) => {
        if (formFields[field as FormFields]?.validations) {
            fields = {
                ...fields,
                [field as FormFields]: '',
            };
        }
    });

    return fields as unknown as Partial<Record<FormFields, string>>;
};

// Check each type of validations and return required error message
const getValidationErrorMessage = (
    formFields: Partial<Record<FormFields, string>>,
    fieldName: FormFields,
    fieldLabel: string,
    validationType: ValidationTypes,
): string => {
    switch (validationType) {
    case 'isRequired':
        return formFields[fieldName] ? '' : `${fieldLabel} is required`;
    default:
        return '';
    }
};

// Genric hook to handle form fields
function useForm(initialFormFields: Partial<Record<FormFields, ModalFormFieldConfig>>) {
    // Form field values
    const [formFields, setFormFields] = useState(getInitialFieldsValue(initialFormFields));

    // Form field error state
    const [errorState, setErrorState] = useState<Partial<Record<FormFields, string>>>(
        getFieldslWhereErrorCheckRequired(initialFormFields),
    );

    /**
     * Set new field value on change
     * and reset field error state
     */
    const onChangeOfFormField = (fieldName: FormFields, value: string) => {
        setErrorState({...errorState, [fieldName]: ''});
        setFormFields({...formFields, [fieldName]: value});
    };

    // Validate all form fields and set error if any
    const isErrorInFormValidation = (): boolean => {
        let fields = {};
        Object.keys(initialFormFields).forEach((field) => {
            if (initialFormFields[field as FormFields]?.validations) {
                Object.keys(initialFormFields[field as FormFields]?.validations ?? '').forEach((validation) => {
                    const validationMessage = getValidationErrorMessage(
                        formFields,
                        field as FormFields,
                        initialFormFields[field as FormFields]?.label || '',
                        validation as ValidationTypes,
                    );
                    if (validationMessage) {
                        fields = {
                            ...fields,
                            [field]: validationMessage,
                        };
                    }
                });
            }
        });

        if (!Object.keys(fields).length) {
            return false;
        }

        setErrorState(fields);
        return true;
    };

    // Reset form field values and error states
    const resetFormFields = () => {
        setFormFields(getInitialFieldsValue(initialFormFields));
        setErrorState(getFieldslWhereErrorCheckRequired(initialFormFields));
    };

    // Set value for a specific form field
    const setSpecificFieldValue = (modifiedFormFields:Partial<Record<FormFields, string>>) => {
        setFormFields(modifiedFormFields);
    };

    return {formFields, errorState, setSpecificFieldValue, onChangeOfFormField, isErrorInFormValidation, resetFormFields};
}

export default useForm;
