#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <asm/uaccess.h>
#include <linux/hugetlb.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/fs.h>



#define BUFSIZE 150

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Monitor de ram");
MODULE_AUTHOR("David González");

struct sysinfo info; 

static int infoMemoria(struct seq_file *arch, void *v){ 

	unsigned long megas=1024*1024;
    unsigned long total = 0;
    unsigned long libre = 0;
	unsigned long unidad = 0;
	si_meminfo(&info);
	unidad = (unsigned long)info.mem_unit;
	total = info.totalram * unidad;
	total = total/megas;
	libre = (info.freeram + info.bufferram) * unidad;
	libre = libre/megas;
	seq_printf(arch, "Carné:\t201610648\n");
	seq_printf(arch, "Nombre:\tDavid González\n");
	seq_printf(arch, "Memoria Total:\t%lu MB\n", total);
	seq_printf(arch, "Memoria Libre:\t%lu MB\n", libre);
	seq_printf(arch, "Memoria usada:\t%lu %%\n", ((total-libre)*100)/total);

	return 0;

        //seq_printf(arch, "Memoria Total: \t %lu MB\n", (info.totalram * 4)/1024);
        //seq_printf(arch, "Memoria Libre: \t %lu MB\n", (info.freeram * 4)/1024);
        //seq_printf(arch, "Memoria usada: \t %lu %%\n", ((info.freeram * 4)/(info.totalram * 4))*100);
        //return 0;
}

static int evento_abrir(struct inode *inode, struct file *file){
        return single_open(file, infoMemoria, NULL);
};

static struct file_operations operaciones ={
        .open=evento_abrir,
        .read = seq_read
};

static int inicio(void)
{
        proc_create("pp", 0, NULL, &operaciones);
        printk(KERN_ALERT "Mi CARNE 201610648\n");
        return 0;
}

static void fin(void)
{
        remove_proc_entry("pp", NULL);
        printk(KERN_ALERT "Sistemas Operativos 1\n");
}

module_init(inicio);
module_exit(fin);
