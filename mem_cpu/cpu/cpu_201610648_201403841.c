#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <asm/uaccess.h>
#include <linux/hugetlb.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/fs.h>
#include <linux/sched/signal.h>
#include <linux/sched.h>

#define BUFSIZE 150

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Monitor de ram");
MODULE_AUTHOR("David González");

static int infoMemoria(struct seq_file *arch, void *v)
{
        struct task_struct *ts;
        struct task_struct *hijos;
        struct list_head *list;
        seq_printf(arch, "************************************\n");
        seq_printf(arch, "*     Carné:   201610648           *\n");
        seq_printf(arch, "*     Nombre:  David González      *\n");
        seq_printf(arch, "*----------------------------------*\n");
        seq_printf(arch, "*     Carné:   201403841           *\n");
        seq_printf(arch, "*     Nombre:  Huriel Gómez        *\n");
        seq_printf(arch, "************************************\n");
        seq_printf(arch, "*            MODULO CPU            *\n");
        seq_printf(arch, "************************************\n");
        

        for_each_process(ts)
        {
                //seq_printf(arch, "\nPID: %d | NOMBRE: %s | ESTADO: %s \n", ts->pid, ts->comm, ts->state);

                if (ts->state == 0)
                {
                        seq_printf(arch, "\nPID: %d | NOMBRE: %s | ESTADO: %s \n", ts->pid, ts->comm, "Task Running");
                }
                else if (ts->state == 1)
                {
                        seq_printf(arch, "\nPID: %d | NOMBRE: %s | ESTADO: %s \n", ts->pid, ts->comm, "Task Interruptible");
                }
                else if (ts->state == 2)
                {
                        seq_printf(arch, "\nPID: %d | NOMBRE: %s | ESTADO: %s \n", ts->pid, ts->comm, "Task Uninterruptible");
                }
                else if (ts->state == 4)
                {
                        seq_printf(arch, "\nPID: %d | NOMBRE: %s | ESTADO: %s \n", ts->pid, ts->comm, "Task Zombie");
                }
                else if (ts->state == 8)
                {
                        seq_printf(arch, "\nPID: %d | NOMBRE: %s | ESTADO: %s \n", ts->pid, ts->comm, "Task Stopped");
                }
                else
                {
                        seq_printf(arch, "\nPID: %d | NOMBRE: %s | ESTADO: %ld \n", ts->pid, ts->comm, ts->state);
                }

                list_for_each(list, &ts->children)
                {
                        hijos = list_entry(list, struct task_struct, sibling);
                        if (hijos->state == 0)
                        {
                                seq_printf(arch, "|____PID: %d | NOMBRE: %s | ESTADO: %s \n", hijos->pid, hijos->comm, "Task Running");
                        }
                        else if (hijos->state == 1)
                        {
                                seq_printf(arch, "|____PID: %d | NOMBRE: %s | ESTADO: %s \n", hijos->pid, hijos->comm, "Task Interruptible");
                        }
                        else if (hijos->state == 2)
                        {
                                seq_printf(arch, "|____PID: %d | NOMBRE: %s | ESTADO: %s \n", hijos->pid, hijos->comm, "Task Uninterruptible");
                        }
                        else if (hijos->state == 4)
                        {
                                seq_printf(arch, "|____PID: %d | NOMBRE: %s | ESTADO: %s \n", hijos->pid, hijos->comm, "Task Zombie");
                        }
                        else if (hijos->state == 8)
                        {
                                seq_printf(arch, "|____PID: %d | NOMBRE: %s | ESTADO: %s \n", hijos->pid, hijos->comm, "Task Stopped");
                        }
                        else
                        {
                                seq_printf(arch, "|____PID: %d | NOMBRE: %s | ESTADO: %ld \n", hijos->pid, hijos->comm, hijos->state);
                        }
                }
        }

        return 0;
}

static int evento_abrir(struct inode *inode, struct file *file)
{
        return single_open(file, infoMemoria, NULL);
};

static struct file_operations operaciones = {
    .open = evento_abrir,
    .read = seq_read};

static int inicio(void)
{
        proc_create("pp", 0, NULL, &operaciones);
        printk(KERN_ALERT "David González - 201610648\nHuriel Gómez - 201403841");
        return 0;
}

static void fin(void)
{
        remove_proc_entry("pp", NULL);
        printk(KERN_ALERT "SISTEMAS OPERATIVOS 1\n");
}

module_init(inicio);
module_exit(fin);