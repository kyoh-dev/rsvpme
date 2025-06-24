import { parseWithZod } from "@conform-to/zod";
import { z } from "zod";

type DefineFormOpts = {
  id: string;
};

type DefineFormResult = {
  formId: string;
  /* Present if the form was submitted */
  formData?: FormData;
};

export const defineForm = async (request: Request, opts: DefineFormOpts) => {
  const result: DefineFormResult = {
    formId: opts.id,
  };

  if (request.method === "POST") {
    const formData = await request.formData();
    if (formData.get("formId") === opts.id) result.formData = formData;
  }

  return result;
};

export const customZodErrorMap: z.ZodErrorMap = (issue, ctx) => {
  if (issue.code === z.ZodIssueCode.invalid_type) {
    if (issue.received === "undefined") {
      return { message: "required" };
    }
  }
  return { message: ctx.defaultError };
};

export const parseFormData = <Schema extends z.AnyZodObject>(
  schema: Schema,
  formData?: FormData,
) => {
  if (!formData) {
    return { status: "no_form" } as const;
  }
  const form = parseWithZod(formData, {
    schema,
    errorMap: customZodErrorMap,
  });

  return form;
};
