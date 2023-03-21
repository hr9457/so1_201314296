#include <linux/module.h>	/* Needed by all modules */
#include <linux/kernel.h>	/* Needed for KERN_INFO */

#include <linux/init.h>		/* Needed for the macros */
#include <linux/proc_fs.h>
#include <asm/uaccess.h>	
#include <linux/seq_file.h>
#include <linux/hugetlb.h>
#include <linux/sched.h>
#include <linux/sched/signal.h>
#include <linux/cred.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo CPU");

struct task_struct *process, *process_son;
struct list_head *son;

unsigned long ramMemory;

static int write_file(struct seq_file *file , void *v)
{
	seq_printf(file, "[");
	seq_printf(file, "{\"Nombre\":\"\",\"UID\":-1,\"PID\":-1,\"Proceso\":\"\",\"Estado\":\"\",\"VmRSS\":-1,\"Hijos\":[]}");
	for_each_process(process){
		//encontrando los procesos padre
		if (process->mm) {
            ramMemory = get_mm_rss(process->mm) << PAGE_SHIFT;
			seq_printf(file, ",{\"Nombre\":\"\",\"UID\":%d,\"PID\":%d,\"Proceso\":\"%s\",\"Estado\":\"%ld\",\"VmRSS\":%lu,\"Hijos\":[",__kuid_val(process->cred->uid), process->pid, process->comm, process->__state,ramMemory/(1024*1024));
        } else {
            seq_printf(file, ",{\"Nombre\":\"\",\"UID\":%d,\"PID\":%d,\"Proceso\":\"%s\",\"Estado\":\"%ld\",\"VmRSS\":%d,\"Hijos\":[",__kuid_val(process->cred->uid), process->pid, process->comm, process->__state,0);
        }
		seq_printf(file, "{\"Nombre\":\"\",\"UID\":-1,\"PID\":-1,\"Proceso\":\"\",\"Estado\":\"\",\"VmRSS\":-1,\"Hijos\":[]}");
		//Encontrando los hijos de cada proceso
		list_for_each(son,&(process->children)){
			process_son = list_entry(son,struct task_struct, sibling);
			if (process_son->mm) {
            	ramMemory = get_mm_rss(process_son->mm) << PAGE_SHIFT;
				seq_printf(file, ",{\"Nombre\":\"\",\"UID\":%d,\"PID\":%d,\"Proceso\":\"%s\",\"Estado\":\"%ld\",\"VmRSS\":%lu,\"Hijos\":[]}", __kuid_val(process_son->cred->uid), process_son->pid, process_son->comm, process_son->__state,ramMemory/(1024*1024));
			} else {
				seq_printf(file, ",{\"Nombre\":\"\",\"UID\":%d,\"PID\":%d,\"Proceso\":\"%s\",\"Estado\":\"%ld\",\"VmRSS\":%d,\"Hijos\":[]}", __kuid_val(process_son->cred->uid), process_son->pid, process_son->comm, process_son->__state,0);
			}
		}
		seq_printf(file, "]}");
	}
	seq_printf(file, "]");
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

static int __init cpu_201314296_init(void)
{
	proc_create("cpu_201314296",0,NULL, &operations);
	printk(KERN_INFO "MOUNT MODULE CPU - HECTOR OROZCO\n");
	return 0;
}

static void __exit cpu_201314296_exit(void)
{
	remove_proc_entry("cpu_201314296",NULL);
	printk(KERN_INFO "RMMOD MODULE CPU - Primer Semestre 2023\n");
}

module_init(cpu_201314296_init);
module_exit(cpu_201314296_exit);