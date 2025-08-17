import { createContext, useState } from "react";

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "~/components/ui/alert-dialog";

export const DialogContext = createContext({
  openModal: (
    // biome-ignore lint/correctness/noUnusedFunctionParameters: ignore
    title: string,
    // biome-ignore lint/correctness/noUnusedFunctionParameters: ignore
    description: string,
    // biome-ignore lint/correctness/noUnusedFunctionParameters: ignore
    onClickCancel: () => void,
    // biome-ignore lint/correctness/noUnusedFunctionParameters: ignore
    onClickYes: () => void,
  ) => {},
});

export const ConfirmDialogProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [context, setContext] = useState({
    title: "",
    description: "",
    onClickCancel: () => {},
    onClickYes: () => {},
  });
  const { title, description, onClickCancel, onClickYes } = context;
  const [open, setOpen] = useState(false);

  const openModal = (
    title: string,
    description: string,
    onClickCancel: () => void,
    onClickYes: () => void,
  ): void => {
    if (!open) {
      setOpen(true);
      setContext({ title, description, onClickCancel, onClickYes });
    }
  };

  return (
    <DialogContext.Provider
      value={{
        openModal,
      }}
    >
      {children}
      <AlertDialog
        open={open}
        onOpenChange={(value: boolean) => {
          return setOpen(value);
        }}
      >
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>{title}</AlertDialogTitle>
            <AlertDialogDescription>{description}</AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel asChild>
              <button type="button" onClick={() => onClickCancel()}>
                Cancel
              </button>
            </AlertDialogCancel>
            <AlertDialogAction asChild>
              <button type="button" onClick={() => onClickYes()}>
                Yes
              </button>
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </DialogContext.Provider>
  );
};
