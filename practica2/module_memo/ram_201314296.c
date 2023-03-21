#include <linux/module.h>	/* Needed by all modules */
#include <linux/kernel.h>	/* Needed for KERN_INFO */

#include <linux/init.h>		/* Needed for the macros */
#include <linux/proc_fs.h>
#include <asm/uaccess.h>	
#include <linux/seq_file.h>
#include <linux/hugetlb.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo Memory");

struct sysinfo inf;


static int write_file(struct seq_file *file , void *v)
{
	long memoria_total;
    long memoria_libre; 
	long memoria_buffer; 
	long memoria_prueba;

	//llenando con la informacion de la memoria
	si_meminfo(&inf);
 	
	memoria_total = (inf.totalram * inf.mem_unit);
    memoria_libre = (inf.freeram * inf.mem_unit); 
	memoria_buffer = (inf.bufferram * inf.mem_unit); 
	memoria_prueba = (inf.totalhigh * inf.mem_unit);
   
	// visualizacion de la memoria libre
	seq_printf(file, "{\"total_memory\": %li, \"free_memory\":%li, \"buffer_memory\":%li }",memoria_total/1024/1024, memoria_libre/1024/1024, memoria_buffer/1024/1024 );
	return 0;
}

static int to_open_file(struct inode *inode,struct file *file)
{
	return single_open(file,write_file,NULL);
}

static struct proc_ops operations = 
{
	.proc_open = to_open_file,
	.proc_read = seq_read
};

static int __init ram_201314296_init(void)
{
	proc_create("ram_201314296",0,NULL, &operations);
	printk(KERN_INFO "MOUNT MODULE MEMORY - 201314296 \n");
	return 0;
}

static void __exit ram_201314296_exit(void)
{
	remove_proc_entry("ram_201314296",NULL);
	printk(KERN_INFO "RMMOD MODULE MEMORY - LAB SO1 - SECCION A\n");
}

module_init(ram_201314296_init);
module_exit(ram_201314296_exit);